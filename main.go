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
	// Get interfaces
	var interfaceName = "eth0"
	var iface, err = net.InterfaceByName(interfaceName)
	if err != nil {
		fmt.Println("Could not get interface ", interfaceName, ": ", err.Error())
		os.Exit(1)
	}
	//fmt.Println("Interface By Name: ", iface)

	/*
		ifaces, err := net.Interfaces()
		if err != nil {
			fmt.Println("Could not get interfaces: ", err.Error())
		}
	*/

	var ip net.IP

	// Loop through interfaces for addresses
	//for _, i := range ifaces {
	addrs, err := iface.Addrs()
	if err != nil {
		fmt.Println("Error getting addresses: ", err.Error())
		os.Exit(1)
	}

	for _, addr := range addrs {
		var iip net.IP
		//fmt.Println("Addr: ", addr)
		switch v := addr.(type) {
		case *net.IPNet:
			iip = v.IP
		case *net.IPAddr:
			iip = v.IP
		}
		if iip.To4() != nil {
			ip = iip
			fmt.Println("Interface: ", iface.Name, "; IP Address: ", iip)
		}
	}

	if ip.To4() != nil {

		// Listen for incoming connections
		l, err := net.Listen(CONN_TYPE, ip.String()+":"+CONN_PORT)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			os.Exit(1)
		}

		// CLose the listener when the application closes
		defer l.Close()
		fmt.Println("Listening on " + ip.String() + ":" + CONN_PORT)
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
	fmt.Println(buf)

	// Send a response back to the person contacting us
	conn.Write([]byte("Message received."))

	// Close the connection when you're done with it
	conn.Close()
}
