package main

import (
	"net"
	"strconv"
)

func handlePing(conn net.Conn) {
	_, err := conn.Write([]byte("$4" + CRLF + PONG + CRLF))
	handleError(err, "Error writing to connection: ")
}

func handleEcho(input []string, conn net.Conn) {
	var echoMessage = input[0]
	var size = strconv.Itoa(len(echoMessage))

	_, err := conn.Write([]byte("$" + size + CRLF + echoMessage + CRLF))
	handleError(err, "Error writing to connection: ")
}

func handleSet(args SetArgs, conn net.Conn, store *kvStore) {
	var key = args.Key
	var value = args.Val
	var expiry = args.Expiry

	err := store.Set(key, value, expiry)

	handleErrorWithReturn(err, "Error writing to key value store: ")
	_, err = conn.Write([]byte("+" + "OK" + CRLF))
	handleError(err, "Error writing to connection: ")
}

func handleGet(input []string, conn net.Conn, store *kvStore) {
	var key = input[0]

	var value = store.Get(key)

	if value == NullBulkString {
		_, err := conn.Write([]byte(NullBulkString))
		handleError(err, "Error writing to connection: ")
	} else {
		var length = len(value)
		_, err := conn.Write([]byte("$" + strconv.Itoa(length) + CRLF + value + CRLF))
		handleError(err, "Error writing to connection: ")
	}
}
