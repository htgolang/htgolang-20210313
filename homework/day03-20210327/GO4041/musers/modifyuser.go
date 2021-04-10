package musers

import (
	"fmt"
)

var users = []map[string]string{
	{
		"id":   "1",
		"name": "arun",
		"addr": "北京",
		"tel":  "xxx1231213",
	},
	{
		"id":   "2",
		"name": "lujianfeng",
		"addr": "北京",
		"tel":  "152xxxxxxxx",
	},
}

func Modifyeuser() {
	var id string
	fmt.Print("请输入用户ID:")
	fmt.Scan(&id)

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
		}else {
			fmt.Println("没有找到",id)
		}
	}

	fmt.Println(users)
}
