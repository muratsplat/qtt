package server

import (
	"io"
	"log"
	"net"

	"github.com/eclipse/paho.mqtt.golang/packets"
	"github.com/muratsplat/qtt/auth"
	"github.com/muratsplat/qtt/session"
)

type Server struct {
	Auth  auth.IAuth
	Conn  net.Conn
	close chan *bool
}

func New(conn net.Conn, auth auth.IAuth) *Server {
	authProvider := auth
	if auth == nil {
		authProvider = AuthDefault
	}
	return &Server{
		Auth:  authProvider,
		Conn:  conn,
		close: make(chan *bool, 1),
	}
}

func (s *Server) Close() error {
	s.close <- new(bool)
	return nil
}

func (s *Server) Run() {

	try := 0
	for {
		packet, err := packets.ReadPacket(s.Conn)
		if err != nil {
			if err == io.EOF {
				s.Conn.Close()
				break
			}

		}

		switch v := packet.(type) {
		case *packets.ConnectPacket:
			try++
			if v.ProtocolName == "MQTT" || v.ProtocolName == "MQIsdp" {
				if v.ProtocolVersion == 4 || v.ProtocolVersion == 3 {
					err := s.Auth.Check(v.Username, string(v.Password))
					if err != nil {
						if err == auth.NotAuthErr {
							unAuth := packets.NewControlPacket(packets.Connack)
							ackPack := unAuth.(*packets.ConnackPacket)
							ackPack.ReturnCode = packets.ErrRefusedNotAuthorised
							err := ackPack.Write(s.Conn)
							if err != nil {
								log.Println(err)
							}
							log.Printf("Client: %s is not authorized. ", v.ClientIdentifier)
							if try >= 2 {
								s.Conn.Close()
							}
						}
						break
					}

					session.Clients.List[v.ClientIdentifier] = session.NewSession(
						v.ClientIdentifier,
						s.Conn,
						s.close,
					)

					ok := packets.NewControlPacket(packets.Connack)
					ackPack := ok.(*packets.ConnackPacket)
					ackPack.ReturnCode = packets.Accepted
					err = ackPack.Write(s.Conn)
					if err != nil {
						log.Println(err)
						log.Println("Connection is closing")
						s.Conn.Close()
					}

					log.Printf("Client: %s is connected. ", v.ClientIdentifier)
					session.Clients.List[v.ClientIdentifier].Run() // blocking

				}
			}

			if session.Clients.List[v.ClientIdentifier].Done {
				return
			}

		default:
			panic("Not handled logic")
		}
	}
}

// TODO:
var AuthDefault auth.IAuth = &auth.Auth{}
