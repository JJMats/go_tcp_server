package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Listen for incoming connections
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// CLose the listener when the application closes
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
		}

		// Handle connections in a new goroutine
		go handleRequest(conn)
	}
}

// Handles incoming requests
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading", err.Error())
	}

	fmt.Println("Request length: ", reqLen)

	// Send a response back to the person contacting us
	conn.Write([]byte("Message received."))

	// Close the connection when you're done with it
	conn.Close()
}
