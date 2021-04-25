package main

import (
	"fmt"
	"strings"
	"usermanage/user"

	// "usermanage/command"
	_ "usermanage/init"
	"usermanage/login"
	"usermanage/tips"
	"usermanage/todo"
)

func main() {
	login.Login()
	users := map[string]map[string]string{}
	var n string
	id := 1
	for {
		fmt.Println(strings.Repeat("*",20))
		fmt.Println("欢迎登录xxx用户系统")
		fmt.Println(
"1. 退出",
)
		for _,v := range tips.Mes {
			fmt.Printf("%s. %s\n", v[0], v[1])
		}
		for _,v := range tips.TaskMes {
			fmt.Printf("%s, %s\n", v[0], v[1])
		}
		fmt.Println(strings.Repeat("*",20))
		fmt.Println("请输入编号:")
		fmt.Scan(&n)

		switch {
		case n == "2":
			id++
			user.UserAdd(id, users)
		case n == "3":
			user.Query(users)
		case n == "4":
			user.DelUser(users)
		case n == "5":
			user.ModifUser(users)
		case n == "6":
			todo.TaskAdd()
		case n == "7":
			todo.TaskDel()
		case n == "8":
			todo.TaskMod()
		case n == "9":
			todo.TaskQuery()
		case n == "1":
			return
		default:
			fmt.Println("指令错误")
		}
	}
}
