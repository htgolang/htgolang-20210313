package user

import (
	"fmt"
)

func (u *Users) generateId() int {
	return len(u.users) + 1
}

func (u *Users) Add() {
	var name, tel, addr string
	fmt.Print("请输入用户名:")
	fmt.Scan(&name)
	fmt.Print("请输入地址:")
	fmt.Scan(&addr)
	fmt.Print("请输入联系方式:")
	fmt.Scan(&tel)

	// 生成id
	id := u.generateId()

	// 添加user
	u.users = append(u.users, user{id, name, addr, tel, true})

	fmt.Println(u.users)
}

// displayUsers 过滤 false 即删除的用户数据
func (u *Users) displayUsers() {
	for _, user := range u.users {
		if user.flage {
			fmt.Println(user.id, user.name, user.addr, user.tel)
		}
	}
}

func Add(u *Users) {
	var name, tel, addr string
	fmt.Print("请输入用户名:")
	fmt.Scan(&name)
	fmt.Print("请输入地址:")
	fmt.Scan(&addr)
	fmt.Print("请输入联系方式:")
	fmt.Scan(&tel)

	// 生成id
	id := u.generateId()

	// 添加user
	u.users = append(u.users, user{id, name, addr, tel, true})

	u.displayUsers()
}
