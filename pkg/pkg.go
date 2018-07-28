package pkg

import (
	"io"

	"github.com/eclipse/paho.mqtt.golang/packets"
)

func Read(r io.Reader) (packets.ConnectPacket, error) {
	return packets.ReadPacket(r)
}
