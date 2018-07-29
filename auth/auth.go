package auth

import (
	"github.com/eclipse/paho.mqtt.golang/packets"
)

var (
	NotAuthErr error = packets.ConnErrors[packets.ErrRefusedBadUsernameOrPassword]
)

type IAuth interface {
	Check(user, pass string) error
}

type Auth struct{}

func (a *Auth) Check(user, pass string) error {
	if user == "user" {
		if pass == "secret" {
			return nil
		}
	}
	return NotAuthErr
}

var _ IAuth = &Auth{}
