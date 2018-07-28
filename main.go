package main

import (
	"net"
)

func main() {

	ln, err := net.Listen("tcp", ":1882")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

}
