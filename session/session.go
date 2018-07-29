package session

import (
	"net"
	"sync"
)

type Session struct {
	Conn     net.Conn
	ClientID string
	Auth     bool
	sync.RWMutex
	Done chan bool
}

type Connections struct {
	sync.RWMutex
	List map[string]*Session
}

func (s *Session) Run() {
	for {

	}
}

var Clients = &Connections{
	List: make(map[string]*Session),
}

func NewSession(clientID string, conn net.Conn) *Session {
	return &Session{
		ClientID: clientID,
		Conn:     conn,
		Auth:     false,
	}
}
