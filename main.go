package main

import (
	"net"

	"github.com/eclipse/paho.mqtt.golang/packets"
	"github.com/muratsplat/qtt/session"
)

func main() {

	ln, err := net.Listen("tcp", ":1883")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// Todo
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	for {
		packet, err := packets.ReadPacket(conn)
		if err != nil {

		}

		switch v := packet.(type) {
		case *packets.ConnectPacket:
			if v.ProtocolName == "MQTT" {
				if v.ProtocolVersion == 4 {
					resp := packets.NewControlPacket(packets.Connack)
					err = resp.Write(conn)
					if err != nil {
						panic(err)
					}
					session.Clients.List[v.ClientIdentifier] = session.NewSession(
						v.ClientIdentifier,
						conn,
					)
					go session.Clients.List[v.ClientIdentifier].Run()

				}
			}

		}
	}

}

var AuthSrv IAuth = &Auth{}
