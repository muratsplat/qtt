package main

// Some references
// - https://github.com/eclipse/mosquitto/blob/master/lib/mqtt3_protocol.h
// - https://github.com/eclipse/paho.mqtt.golang/blob/master/packets/packets.go#L58

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/muratsplat/qtt/auth"
	"github.com/muratsplat/qtt/conf"
	"github.com/muratsplat/qtt/server"
)

var (
	shutdown    = make(chan os.Signal, 1)
	Listener    net.Listener
	defaultAuth = &auth.Auth{
		User: conf.Get().DefaultUser,
		Pass: conf.Get().DefaultPass,
	}
)

func init() {
	signal.Notify(shutdown, os.Interrupt)
	var err error
	addr := fmt.Sprintf(":%s", conf.Get().MQTTPort)
	Listener, err = net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	log.Printf("%s address is listening...\n", addr)
}

func main() {

	var closing bool
	go func() {
		<-shutdown
		closing = true
		time.Sleep(time.Second * 5)
		log.Println("Server is shutdown...")
		Listener.Close()
	}()
	for {
		conn, err := Listener.Accept()

		if err != nil {
			if closing {
				break
			}
			panic(err)
		}
		srv := server.New(conn, defaultAuth)
		go srv.Run()
	}
}
