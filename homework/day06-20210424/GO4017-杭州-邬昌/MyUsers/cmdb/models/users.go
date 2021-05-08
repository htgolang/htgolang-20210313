package models

import "fmt"

type User struct {
	Id   string
	Name string
	Addr string
	Tel  string
}

type UserList struct {
	Userlist map[string]*User
}

func NewUser(id, name, addr, tel string) *User {
	return &User{
		Id:   id,
		Name: name,
		Addr: addr,
		Tel:  tel,
	}
}

//var Us = make(map[string]*User, 40)

func (u *User) String() string {

	return fmt.Sprintf("用户id: %v\n用户名: %v\n地址: %v\n电话: %v\n", u.Id, u.Name, u.Addr, u.Tel)
}
