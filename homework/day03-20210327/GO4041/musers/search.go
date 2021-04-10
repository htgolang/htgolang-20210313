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

func Searcheuser() {
	var id string
	fmt.Print("请输入用户ID:")
	fmt.Scan(&id)

	for _, user := range users {
		if id == user["id"] {
			fmt.Println(user)
		}
	}

	fmt.Println(users)
}
