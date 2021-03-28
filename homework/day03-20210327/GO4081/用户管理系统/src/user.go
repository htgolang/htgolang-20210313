/**
* @Author: fengzhilaoling
* @Date: 2021/3/28 9:28
* 用户管理系统
**/
package main

import "fmt"

func main() {
	users := []map[string]string{
		{
			"id":      "1",
			"name":    "老柯",
			"address": "北京",
			"tel":     "15287869909",
		},
		{
			"id":      "2",
			"name":    "小胡",
			"address": "北京",
			"tel":     "13444636678",
		},
	}
	//fmt.Println(users)
	userSystem(users)

}

func userSystem(users []map[string]string) {
	fmt.Println("--------------------欢迎来到用户管理系统-------------------")
	for {
		fmt.Println("1. 查看系统所有用户")
		fmt.Println("2. 添加用户")
		fmt.Println("3. 删除用户")
		fmt.Println("4. 修改用户信息")
		fmt.Println("5. 查询用户信息")
		fmt.Println("6. 退出系统")
		var codeNum int
		fmt.Print("请输入你要操作的代码：")
		fmt.Scanln(&codeNum)
		switch codeNum {
		case 1:
			// 查看所有用户信息
			viewUser(users)
		case 2:
			// 添加用户
			users = addUser(users)
		case 3:
			// 删除用户
			users = delUser(users)
		case 4:
			// 修改用户信息
			users = updateUser(users)
		case 5:
			// 查询用户信息
			getUser(users)
		case 6:
			// 退出系统
			fmt.Println("系统退出成功！")
			break
		default:
			fmt.Println("输入的操作代码不正确！！！")
		}
	}

}
