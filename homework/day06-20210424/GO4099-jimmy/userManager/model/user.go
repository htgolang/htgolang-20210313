package model

import (
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

func (u User) String() string {
	fmt.Println(strings.Repeat("=", 20))
	return fmt.Sprintf("用户id: %d\n用户名: %v\n地址: %v\n电话: %v\n密码: %v\n", u.Id, u.Name, u.Addr, u.Tel, u.Password)
}

func NewUser(id int, name, addr, tel, pwd string) User {
	return User{
		Id:       id,
		Name:     name,
		Addr:     addr,
		Tel:      tel,
		Password: pwd,
	}
}
