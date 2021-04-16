package main

import (

	"cmdb/users"
	"fmt"
	"github.com/princebot/getpass"
)

func main() {

	var flag int
	const pass = "12345"

	/*
		执行操作：cd
		1、添加用户  2、删除用户
		3、查询用户  4、修改用户
	*/
	fmt.Println("------欢迎进入用户管理系统-----，请输入用户密码：")
	for {

		//fmt.Scan(&password)
		password, _ := getpass.Get("请输入密码：")
		if string(password) == pass {

			fmt.Println("欢迎，请继续操作！！！")
			break

		} else {

			flag++
			fmt.Println("密码错误，请重新输入密码！！！")

			if flag > 2 {

				fmt.Println("密码错误3次了，不能再尝试！！！！")
				return
			}
		}
	}


	for {

		users.Prom()
		cmd :=users.Input("请输入指令：")

		switch cmd {

		case "1":
			fmt.Println(users.AddUser())
		case "2":
			fmt.Println(users.DeletUser())
		case "3":
			fmt.Println(users.QueryUsers())
		case "4":
			fmt.Println(users.ModifyUsers())
		default:
			fmt.Println("退出游戏！！！")

			return
		}

	}

}