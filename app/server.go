package main

import (
	"net"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:6379")

	handleErrorWithExit(err, "Failed to bind to port 6379")

	defer l.Close()

	eventLoop := NewEventLoop()

	gomuKvStore := createStore()

	eventLoop.Start(gomuKvStore)

	for {
		conn, err := l.Accept()

		handleErrorWithExit(err, "Error accepting connection: ")

		eventLoop.read(conn)
	}
}
