package model

import (
	"fmt"
	"strings"
)

type User struct {
	Id   int
	Name string
	Addr string
	Tel  string
}

type UserList struct {
	Users map[int]User
}

func NewUser(id int, name, addr, tel string) *User {
	return &User{
		Id:   id,
		Name: name,
		Addr: addr,
		Tel:  tel,
	}
}

func (u User) String() string {
	fmt.Println(strings.Repeat("=", 20))
	return fmt.Sprintf("用户id: %d\n用户名: %v\n地址: %v\n电话: %v\n", u.Id, u.Name, u.Addr, u.Tel)
}
