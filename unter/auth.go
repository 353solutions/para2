package main

type User struct {
	Login string
}

type Auth struct {
}

func NewAuth() *Auth {
	a := Auth{}
	return &a
}

func (a *Auth) Login(user, passwd string) (User, bool) {
	if user == "joe" && passwd == "baz00ka" {
		return User{"joe"}, true
	}

	return User{}, false
}
