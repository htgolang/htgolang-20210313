package main

import (
	"fmt"
	"strings"
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

func searchUser() {
	var s string
	fmt.Print("请输入查询字符串:")
	fmt.Scanf("%s", &s)
	for _, user := range users {
		if strings.Contains(user["id"], s) || strings.Contains(user["name"], s) || strings.Contains(user["address"], s) {
		    fmt.Println("查询成功!")
		    fmt.Printf("id: %s\n", user["id"])
		    fmt.Printf("name: %s\n", user["name"])
		    fmt.Printf("address: %s\n", user["address"])
			fmt.Println()
		}
	}
}

func main() {
	searchUser()
	fmt.Println(users)
}