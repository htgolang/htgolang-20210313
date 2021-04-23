package user

import (
	"fmt"
	"strconv"
)

func (u *Users) Modify() {
	var id string
	fmt.Print("请输入用户ID:")
	fmt.Scan(&id)
	idn, _ := strconv.Atoi(id)

	status := false

	for _, usr := range u.users {
		if idn == usr.id {
			var name, tel, addr string
			fmt.Print("请输入用户名:")
			fmt.Scan(&name)
			fmt.Print("请输入地址:")
			fmt.Scan(&addr)
			fmt.Print("请输入联系方式:")
			fmt.Scan(&tel)
			usr = user{idn, name, tel, addr, true}
		}
	}

	if !status {
		fmt.Println("无该用户信息！")
	}

	u.displayUsers()
}

func Modify(u *Users) {
	var id string
	fmt.Print("请输入用户ID:")
	fmt.Scan(&id)
	idn, _ := strconv.Atoi(id)

	status := false

	for i, usr := range u.users {
		if idn == usr.id {
			var name, tel, addr string
			usr.displayUser()
			fmt.Print("请输入用户名:")
			fmt.Scan(&name)
			fmt.Print("请输入地址:")
			fmt.Scan(&addr)
			fmt.Print("请输入联系方式:")
			fmt.Scan(&tel)
			usr = user{idn, name, tel, addr, true}
			u.users[i] = usr

			status = true
		}
	}

	if !status {
		fmt.Println("无该用户信息！")
	}

	u.displayUsers()
}

// displayUser 过滤 false 即删除的用户数据
func (u *user) displayUser() {
	fmt.Println(u.id,u.name, u.addr, u.tel, )
}
