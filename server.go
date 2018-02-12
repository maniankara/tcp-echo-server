package main
/*
	A simple TCP echo server on the current host with the given port
	Replies with the same data with host information

	License: Apache License-2.0

	*/


import (
	"fmt"
	"flag"
	"net"
	"os"
)

/*
	Handling and printing error

	*/

func handleError(err error) {
	// handle error
	if err != nil {
		fmt.Println("Error: ", err)
	}	
}

/* 
	Start 

	*/
func main() {

	var port string

	/* command line */
	flag.StringVar(&port, "port", "8445", "Port where server listens to")
	flag.Parse()

	/* TCP listener */
	ln, err := net.Listen("tcp", ":" + port)
	handleError(err)
	// create the connection
	for {
		conn, err := ln.Accept()
		handleError(err)

		// handle the connection
		go handlePackets(conn)
	}

}

// ref: https://github.com/golergka/go-tcp-echo/blob/master/go-tcp-echo.go
/*
	Reply with whatever received with the current hostname

	*/
func handlePackets(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("Closing connection")

	hostname, err := os.Hostname()
	handleError(err)
	hostname += ": "

	for {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		if err != nil {
			handleError(err)
			return
		}
		data := buf[:size]
		host := []byte(hostname)
		hostData := append(host, data...)
		conn.Write(hostData)
	}	
	
}