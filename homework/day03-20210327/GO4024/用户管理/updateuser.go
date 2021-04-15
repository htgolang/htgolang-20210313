package main

import (
	"fmt"
	_ "strconv"
	_ "strings"
)

var users = []map[string]string {
	{
		"id":"1",
		"name":"arun",
		"addr": "Beijing",
		"tel": "xxx123123123",
	},
	{
		"id" : "2",
		"name": "ujianfeng",
		"addr":"Beijing",
		"tel":"152xxxxxxx",
	},
}


//修改用户信息
func modifUser() {
	var nname,orginame,addrs,phones string
	fmt.Println("请输入要修改用户名:")
	fmt.Scan(&orginame)
	for k,_ := range users {
		if orginame == users[k]["name"] {
			fmt.Println("请输入新的用户名:")
			fmt.Scan(&nname)
			fmt.Println("请输入新地址:")
			fmt.Scan(&addrs)
			fmt.Println("请输入新电话:")
			fmt.Scan(&phones)
			users[k]["name"] = nname
			users[k]["addr"] = addrs
			users[k]["tel"] = phones
		} 
	}
	fmt.Println(users)
}

func main() {
	modifUser()

}