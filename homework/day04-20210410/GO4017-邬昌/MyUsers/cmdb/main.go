package main

import (
	"cmdb/handles"
	"cmdb/users"
	"fmt"
	_ "cmdb/init"
)

func main() {

	/*
		执行操作：cd
		1、添加用户  2、删除用户
		3、查询用户  4、修改用户
	*/
	fmt.Println("------欢迎进入用户管理系统-----，请输入用户密码：")
	handles.CheckPassword()


	for {

		//commands.Prom
		handles.Prompt()
		cmd :=users.Input("请输入指令：")

		if comm,ok :=handles.Commands[cmd]; ok{

			comm()
		}else if cmd =="1"{

			break
		}else {

			fmt.Println("指令错误")
		}


	}

}