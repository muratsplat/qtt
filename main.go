package main

import (
	"fmt"
	"net"

	"github.com/eclipse/paho.mqtt.golang/packets"
)

func main() {

	ln, err := net.Listen("tcp", ":1883")
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

	for {

		packet, err := packets.ReadPacket(conn)
		if err != nil {
			panic(err)
		}

		switch v := packet.(type) {
		case *packets.ConnectPacket:
			pass := string(v.Password)
			user := v.Username
			err := AuthSrv.Check(user, pass)
			if err != nil {

			}

		}

		fmt.Println(packet.String())

	}

}

var AuthSrv IAuth = &Auth{}
