package main

import (
	"fmt"
)

var users = []map[string]string{
	{
		"id": "1",
		"name": "foo",
		"address": "beijing",
	},
	{
		"id": "2",
		"name": "bar",
		"address": "shanghai",
	},
	{
		"id": "3",
		"name": "foobar",
		"address": "guangzhou",
	},
}


func modifyUser() {
	var id, name, address string
	fmt.Print("请输入要修改的用户id:")
	fmt.Scanf("%s", &id)
	var i int
	for i=0;i<len(users);i++ {
		if users[i]["id"] == id {
			break
		}
	}
	if i == len(users) {
		fmt.Println("用户不存在!")
		return
	}
	fmt.Print("请输入新的用户名:")
	fmt.Scanf("%s", &name)
	fmt.Print("请输入新的地址:")
	fmt.Scanf("%s", &address)
	user := map[string]string{"id": id, "name": name, "address": address}
	users[i] = user
}

func main() {
	modifyUser()
	fmt.Println(users)
}