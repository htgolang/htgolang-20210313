package user

import "fmt"

func Modify() {
	var id string
	fmt.Print("请输入用户ID:")
	fmt.Scan(&id)

	status := false

	for _, user := range users {
		if id == user["id"] {
			fmt.Println(user)
			var name, tel, addr string
			fmt.Print("请输入用户名:")
			fmt.Scan(&name)
			fmt.Print("请输入地址:")
			fmt.Scan(&addr)
			fmt.Print("请输入联系方式:")
			fmt.Scan(&tel)

			user["id"] = id
			user["name"] = name
			user["addr"] = addr
			user["tel"] = tel

			status = true
		}
	}

	if !status {
		fmt.Println("无该用户信息！")
	}

	fmt.Println(users)
}
