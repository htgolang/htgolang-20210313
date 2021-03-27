package main

import (
	"fmt"
	"strconv"
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

func main() {

	var name, tel, addr string
	fmt.Print("请输入用户名:")
	fmt.Scan(&name)
	fmt.Print("请输入地址:")
	fmt.Scan(&addr)
	fmt.Print("请输入联系方式:")
	fmt.Scan(&tel)

	id := 0
	for _, user := range users {
		i, _ := strconv.Atoi(user["id"])
		if i > id {
			id = i
		}
	}
	users = append(users, map[string]string{
		"id":   strconv.Itoa(id + 1),
		"name": name,
		"addr": addr,
		"tel":  tel,
	})

	fmt.Println(users)
}
