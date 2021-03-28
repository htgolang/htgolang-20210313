/**
* @Author: fengzhilaoling
* @Date: 2021/3/28 12:15
**/
package main

import "fmt"

func getUser(users []map[string]string) {
	var str1 string
	fmt.Print("请输入要查询用户的姓名或地址或手机号：")
	fmt.Scanln(&str1)
	for _, value := range users {
		switch {
		case str1 == value["name"]:
			fmt.Printf("用户的ID：%v  用户的姓名：%v  用户的地址：%v, 用户的手机号：%v\n", value["id"],
				value["name"], value["address"], value["tel"])
			// 存在用户名相同的情况
			continue
		case str1 == value["address"]:
			fmt.Printf("用户的ID：%v  用户的姓名：%v  用户的地址：%v, 用户的手机号：%v\n", value["id"],
				value["name"], value["address"], value["tel"])
			// 存在用户地址相同的情况
			continue
		case str1 == value["tel"]:
			fmt.Printf("用户的ID：%v  用户的姓名：%v  用户的地址：%v, 用户的手机号：%v\n", value["id"],
				value["name"], value["address"], value["tel"])
		}
	}
	fmt.Println("你要查询的用户不存在！")
}
