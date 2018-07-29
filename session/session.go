package session

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"github.com/eclipse/paho.mqtt.golang/packets"
	"github.com/muratsplat/qtt/auth"
)

type Session struct {
	Conn     net.Conn
	ClientID string
	Auth     bool
	sync.RWMutex
	Done       bool
	User, Pass string
}

type Connections struct {
	sync.RWMutex
	List map[string]*Session
	Auth auth.IAuth
}

func (s *Session) Run() {
	for {

		packet, err := packets.ReadPacket(s.Conn)
		if err != nil {
			if err == io.EOF {
				log.Printf("It looks clinet %s closed the connection \n", s.ClientID)
				s.Conn.Close()
				s.Done = true
				break
			}
		}

		switch v := packet.(type) {
		case *packets.PingreqPacket:
			pong := packets.NewControlPacket(packets.Pingresp)
			pong.Write(s.Conn)
			log.Printf("Sending pong request for  (%s) the connection \n", s.ClientID)
			_ = v

		}

		fmt.Println(packet.String())

	}
}

var Clients = &Connections{
	List: make(map[string]*Session),
	Auth: &auth.Auth{},
}

func NewSession(clientID string, conn net.Conn) *Session {
	return &Session{
		ClientID: clientID,
		Conn:     conn,
		Auth:     false,
		Done:     false,
	}
}
