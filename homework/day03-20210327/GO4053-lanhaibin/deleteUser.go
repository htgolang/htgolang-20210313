package main

import (
	"fmt"
	"strconv"
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

func deleteUser() {
	var id string
	fmt.Print("请输入要删除的用户id:")
	fmt.Scanf("%s", &id)
	var i int
	for i=0;i<len(users);i++ {
		if users[i]["id"] == id {
			break
		}
	}
	if i != len(users) {
		users = append(users[:i], users[i+1:]...)
		fmt.Println("删除成功")
	} else {
		fmt.Println("为找到该用户")
	}
}

func main() {
	deleteUser()
	fmt.Println(users)
}