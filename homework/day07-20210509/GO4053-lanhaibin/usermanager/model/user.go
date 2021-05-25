package model

import (
	"crypto/sha512"
	"fmt"
	"strings"
)

type User struct {
	Id       int
	Name     string
	Addr     string
	Tel      string
	Password string
}

type UserList []*User

var Users map[int]*User = map[int]*User{
	1: &User{
		Id:       1,
		Name:     "legion",
		Password: "ba3253876aed6bc22d4a6ff53d8406c6ad864195ed144ab5c87621b6c233b548baeae6956df346ec8c17f5ea10f35ee3cbc514797ed7ddd3145464e2a0bab413",
	},
}

func GetUserId() int {
	max := 1
	for i, _ := range Users {
		if i > max {
			max = i
		}
	}
	return max + 1
}

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

func (u User) CheckPassword(password string) bool {
	hash := sha512.New()
	hash.Write([]byte(password))
	return u.Password == fmt.Sprintf("%x", string(hash.Sum([]byte{})))
}
