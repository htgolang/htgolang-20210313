package user

import "fmt"

func Delete() {
	var id string
	fmt.Print("请输入用户ID:")
	fmt.Scan(&id)

	status := false

	for i, user := range users {
		if id == user["id"] {
			fmt.Println(user)
			var rep string
			fmt.Print("请确认是否删除:")
			fmt.Scan(&rep)
			if rep == "y" || rep == "Y" {
				users = append(users[:i], users[i+1:]...)
			}
			break
			status = true
		}
	}

	if !status {
		fmt.Println("无该用户信息！")
	}

	fmt.Println(users)
}
