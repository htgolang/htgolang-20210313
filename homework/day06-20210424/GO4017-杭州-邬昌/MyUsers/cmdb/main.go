package main

import (
	"cmdb/handles"
	_ "cmdb/init"
	"cmdb/utils"
	"fmt"
	"strconv"
	//"github.com/howeyc/gopass"
)

func main() {

	/*
		执行操作：cd
		1、添加用户  2、删除用户
		3、查询用户  4、修改用户
	*/
	fmt.Println("------欢迎进入用户管理系统-----")
	ckpassword := handles.CheckPassword()
	if ckpassword == false {

		return
	}
	//user1 := models.NewUser("1", "小明", "honkong", "18992928820")
	//fmt.Println(user1.String())

	for {

		//commands.Prom
		handles.Prompt()
		cmd := utils.Input("请输入指令：")

		if cmd, err := strconv.Atoi(cmd); err == nil {

			if commands := handles.Tip.GetID(cmd); commands != nil {

				commands()
			}

		} else if cmd == 1 {

			break

		} else {

			fmt.Println("指令错误")
		}

	}

}
