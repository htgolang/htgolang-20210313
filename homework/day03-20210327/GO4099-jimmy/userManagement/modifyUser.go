package main

import (
	"fmt"
	"strconv"
)

func modifyUser() {
	var id int
	var confirm string
	fmt.Print("请输入要修改的用户ID: ")
	fmt.Scan(&id)
	for i, v := range users {
		uid, err := strconv.Atoi(v["id"])
		if err != nil {
			fmt.Println(err)
			return
		}
		if uid == id {
			fmt.Printf("id为%d的用户存在, 用户信息: %v\n", id, v)
			fmt.Print("请用户确认修改(y/yes/Y/YES): ")
			fmt.Scan(&confirm)
			switch confirm {
			case "y", "yes", "Y", "YES":
				var name, addr, tel string
				fmt.Print("请输入用户名: ")
				fmt.Scan(&name)
				fmt.Print("请输入地址: ")
				fmt.Scan(&addr)
				fmt.Print("请输入电话: ")
				fmt.Scan(&tel)
				users[i]["name"] = name
				users[i]["addr"] = addr
				users[i]["tel"] = tel
				fmt.Println("修改以后用户的信息为: ", users[i])
				return
			default:
				fmt.Println("未修改")
				return
			}
		}
	}
	fmt.Println("无此用户")
}
