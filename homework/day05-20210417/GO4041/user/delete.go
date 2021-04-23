package user

import (
	"fmt"
	"strconv"
)

func (u *Users) Delete() {
	var id string
	fmt.Print("请输入用户ID:")
	fmt.Scan(&id)
	idn, _ := strconv.Atoi(id)

	status := false

	for _, user := range u.users {
		if idn == user.id {
			fmt.Println(user)
			var rep string
			fmt.Print("请确认是否删除:")
			fmt.Scan(&rep)
			if rep == "y" || rep == "Y" || rep == "yes" || rep == "Yes" {
				user.flage = false
			}
			status = true
			break
		}
	}

	if !status {
		fmt.Println("无该用户信息！")
	}

	u.displayUsers()
}

func Delete(u *Users) {
	var id string
	fmt.Print("请输入用户ID:")
	fmt.Scan(&id)
	idn, _ := strconv.Atoi(id)

	status := false

	for i, user := range u.users {
		if idn == user.id {
			user.displayUser()
			var rep string
			fmt.Print("请确认是否删除:")
			fmt.Scan(&rep)
			if rep == "y" || rep == "Y" || rep == "yes" || rep == "Yes" {
				user.flage = false
			}
			status = true
			u.users[i] = user
			break
		}
	}

	if !status {
		fmt.Println("无该用户信息！")
	}

	u.displayUsers()
}
