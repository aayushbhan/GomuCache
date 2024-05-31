package main

import (
	"fmt"
	"net"
	"sync"
)

const (
	PING           = "PING"
	PONG           = "PONG"
	ECHO           = "ECHO"
	SET            = "SET"
	GET            = "GET"
	CRLF           = "\r\n"
	NullBulkString = "$-1" + CRLF
)

const ReadBufferSize = 128

type EventLoop struct {
	WorkerPoolSize uint8
	WorkerQueue    chan net.Conn
	WaitGroup      sync.WaitGroup
}

func NewEventLoop() *EventLoop {
	return &EventLoop{
		WorkerPoolSize: 5,
		WorkerQueue:    make(chan net.Conn, 5),
	}
}

func (el *EventLoop) Start(kvStore *kvStore) {
	// Initialize worker pool
	for i := uint8(0); i < el.WorkerPoolSize; i++ {
		el.WaitGroup.Add(1)
		go worker(el, kvStore)
	}
}

func (el *EventLoop) read(conn net.Conn) {
	el.WorkerQueue <- conn
}

func worker(el *EventLoop, store *kvStore) {
	defer el.WaitGroup.Done()

	for conn := range el.WorkerQueue {
		handleConnection(conn, store)
	}
}

func handleConnection(conn net.Conn, store *kvStore) {
	defer conn.Close()

	readBuffer := make([]byte, ReadBufferSize)
	for {
		n, err := conn.Read(readBuffer)
		if err != nil {
			fmt.Println("Error reading from connection: ", err.Error())
			return
		}

		handleCommand(readBuffer[:n], conn, store)
	}
}

func handleCommand(readBuffer []byte, conn net.Conn, store *kvStore) {
	command := parseRESP(readBuffer)

	switch command.Command {
	case PING:
		handlePing(conn)
	case ECHO:
		handleEcho(command.EchoArgs, conn)
	case SET:
		handleSet(command.SetArgs, conn, store)
	case GET:
		handleGet(command.GetArgs, conn, store)
	}
}
