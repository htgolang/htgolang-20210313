package model

import (
	"fmt"
	"strings"
)

var UserId int = 1

type User struct {
	Id   int
	Name string
	Addr string
	Tel  string
}

type UserList []*User

var Users map[int]*User = make(map[int]*User)

// func NewUser(name, addr, tel string) *User {
// 	user := &User{
// 		Id:   userId,
// 		Name: name,
// 		Addr: addr,
// 		Tel:  tel,
// 	}
// 	userId++
// 	Users[userId] = user
// 	return user
// }

func (u User) String() string {
	return fmt.Sprintf("id: %d, 名称: %s, 地址: %s, 电话: %s\n", u.Id, u.Name, u.Addr, u.Tel)
}

func (u User) Contains(s string) bool {
	if strings.Contains(u.Name, s) {
		return true
	}
	if strings.Contains(u.Addr, s) {
		return true
	}
	if strings.Contains(u.Tel, s) {
		return true
	}
	return false
}
