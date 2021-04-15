package main

import (
	"fmt"
	"strconv"
)
//添加用户
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

func genId() int{
	id := 0
	for _,user := range users {
		i, _ := strconv.Atoi(user["id"])
		if i > id {
			id = i
		}
	}
	return id + 1
}

func input(prompt string) string {
	fmt.Println(prompt)
	var text string
	fmt.Scan(&text)
	return text
}

func adduser() error {
	name := input("请输入用户名:")
	addr := input("请输入地址:")
	tel := input("请输入电话:")

	//验证用户名不能重复

	for k,_ := range users {
		if name == users[k]["name"] {
			return fmt.Errorf("用户名%s已存在",name)
		}
	}
	// if name == "kk" {
	// 	return fmt.Errorf("用户名%s已存在",name)
	// }
	users = append(users,map[string]string{
		"id": strconv.Itoa(genId()),
		"name" : name,
		"addr" : addr,
		"tel" : tel,
	})
	return nil 
}

func main() {

	// var name, tel, addr string
	// fmt.Println("请输入用户名:")
	// fmt.Scan(&name)
	// fmt.Println("请输入地址:")
	// fmt.Scan(&addr)
	// fmt.Println("请输入电话:")
	// fmt.Scan(&tel)
	if err := adduser(); err == nil {
		fmt.Println(users)
	} else {
		fmt.Println("添加失败",err)
	}
	
}