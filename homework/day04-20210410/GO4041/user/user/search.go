package user

import "fmt"

func Search() {
	var id string
	fmt.Print("请输入用户ID:")
	fmt.Scan(&id)

	status := false

	for _, user := range users {
		if id == user["id"] {
			fmt.Println(user)
			status = true
		}
	}
	if !status {
		fmt.Println("没有找到用户信息")
	}

	fmt.Println(users)
}
