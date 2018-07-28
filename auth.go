package main

type IAuth interface {
	Check(user, pass string) error
}

type Auth struct{}

func (a *Auth) Check(user, pass string) error {

	// todo:
	return nil
}

var _ IAuth = &Auth{}
