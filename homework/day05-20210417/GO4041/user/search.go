package user

import (
	"fmt"
	"strconv"
)

func (u *Users) Search() {
	var id string
	fmt.Print("请输入用户ID:")
	fmt.Scan(&id)
	idn, _ := strconv.Atoi(id)

	status := false

	for _, usr := range u.users {
		if idn == usr.id {
			fmt.Println(usr)
			status = true
		}
	}
	if !status {
		fmt.Println("没有找到用户信息")
	}
	u.displayUsers()
}

func Search(u *Users) {
	var id string
	fmt.Print("请输入用户ID:")
	fmt.Scan(&id)
	idn, _ := strconv.Atoi(id)
	status := false
	for _, usr := range u.users {
		if idn == usr.id && usr.flage {
			usr.displayUser()
			status = true
			break
		}
	}
	if !status {
		fmt.Println("没有找到用户信息")
	}
}
