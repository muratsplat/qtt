package session

import (
	"fmt"
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
	Done       chan bool
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
			panic(err)
		}

		fmt.Println(packet.String())

		// // Auth Handling
		// if s.Auth == false {
		// 	err := Clients.Auth.Check(s.User, s.Pass)
		// 	if err != nil {
		// 		unAuth := packets.NewControlPacket(packets.Connack)
		// 		ackPack := unAuth.(*packets.ConnackPacket)
		// 		ackPack.ReturnCode = packets.ErrRefusedBadUsernameOrPassword
		// 		err := ackPack.Write(s.Conn)
		// 		if err != nil {
		// 			log.Println(err)
		// 			log.Println("Connection is closing")
		// 			s.Conn.Close()
		// 		}
		// 	}

		// }
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
	}
}
