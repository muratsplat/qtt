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

type Auth struct {
	User string
	Pass string
}

func (a *Auth) Check(user, pass string) error {

	// Allow anonymous user if any credenials are not set
	if len(a.User) == 0 && len(a.User) == 0 {
		return nil
	}

	if user == a.User {
		if pass == a.Pass {
			return nil
		}
	}
	return NotAuthErr
}
