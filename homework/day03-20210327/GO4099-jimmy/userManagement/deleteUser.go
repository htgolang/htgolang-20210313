package main

import (
	"fmt"
	"strconv"
)

func deleteUser() {
	var id int
	var confirm string
	fmt.Print("请输入要删除的用户ID: ")
	fmt.Scan(&id)
	for i, v := range users {
		uid, err := strconv.Atoi(v["id"])
		if err != nil {
			fmt.Println(err)
			return
		}
		if uid == id {
			fmt.Printf("id为%d的用户存在, 用户信息: %v\n", id, v)
			fmt.Print("请确认删除(y/yes/Y/YES): ")
			fmt.Scan(&confirm)
			switch confirm {
			case "y", "yes", "Y", "YES":
				copy(users[i:], users[i+1:])
				users = users[:len(users)-1]
				fmt.Println(users)
				return
			default:
				fmt.Println("未删除")
				return
			}
		}
	}
	fmt.Println("无此用户")
}