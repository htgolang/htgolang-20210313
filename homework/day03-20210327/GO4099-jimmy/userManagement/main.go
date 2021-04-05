package main

import (
	"fmt"
	"os"
)

func main() {
	var text string
	for {
		fmt.Println("请选择要执行的操作:\n\t添加用户(add/a)\n\t删除用户(delete/d)\n\t修改用户(modify/m)\n\t查询用户(query/q)")
		fmt.Scan(&text)

		switch text {
		case "add", "a":
			addUser()
		case "delete", "d":
			deleteUser()
		case "modify", "m":
			modifyUser()
		case "query", "q":
			queryUser()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}
