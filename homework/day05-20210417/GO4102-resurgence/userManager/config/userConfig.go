package config

import "strconv"

type User struct {
	id          int
	Name, Phone string
}

func NewUser(name, phone string) (*User, string) {
	Uid++
	return &User{
		id:    Uid,
		Name:  name,
		Phone: phone,
	}, strconv.Itoa(Uid)
}

type fu func(map[string]*User)

var (
	Uid          = 0
	UserMap     = make(map[string]*User)
	UserFuncMap = make(map[string]fu)
	UserMenuMap = make(map[string]string)
)


func UserRegister(desc string, callback fu) {
	idx := strconv.Itoa(Uid)
	UserMenuMap[idx] = desc
	UserFuncMap[idx] = callback
	Uid++
}