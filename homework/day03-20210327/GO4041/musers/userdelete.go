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

func Deleteuser() {
	var id string
	fmt.Print("请输入用户ID:")
	fmt.Scan(&id)

	for i, user := range users {
		if id == user["id"] {
			fmt.Println(user)
			var rep string
			fmt.Print("请确认是否删除:")
			fmt.Scan(&rep)
			if rep == "y" || rep == "Y" {
				users = append(users[:i], users[i+1:]...)
			}
		} else {
			fmt.Println("没有找到这个", id)
		}
	}

	fmt.Println(users)
}
