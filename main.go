package main

// Some references
// - https://github.com/eclipse/mosquitto/blob/master/lib/mqtt3_protocol.h
// - https://github.com/eclipse/paho.mqtt.golang/blob/master/packets/packets.go#L58

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/eclipse/paho.mqtt.golang/packets"
	"github.com/muratsplat/qtt/auth"
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
	try := 0
	for {
		packet, err := packets.ReadPacket(conn)
		if err != nil {
			if err == io.EOF {
				conn.Close()
			}

			panic(err)
		}

		switch v := packet.(type) {
		case *packets.ConnectPacket:
			try++
			if v.ProtocolName == "MQTT" || v.ProtocolName == "MQIsdp" {
				if v.ProtocolVersion == 4 || v.ProtocolVersion == 3 {
					err := Auth.Check(v.Username, string(v.Password))
					if err != nil {
						if err == auth.NotAuthErr {
							unAuth := packets.NewControlPacket(packets.Connack)
							ackPack := unAuth.(*packets.ConnackPacket)
							ackPack.ReturnCode = packets.ErrRefusedNotAuthorised
							fmt.Println(ackPack.String())
							err := ackPack.Write(conn)
							if err != nil {
								log.Println(err)
							}
							if try >= 1 {
								log.Println("Connection is closing")
								conn.Close()
								break
							}
						}

						log.Println("Connection is closing")
						conn.Close()

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

}

var Auth auth.IAuth = &auth.Auth{}
