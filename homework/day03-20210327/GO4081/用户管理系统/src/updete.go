/**
* @Author: fengzhilaoling
* @Date: 2021/3/28 12:00
* 修改用户信息
**/
package main

import "fmt"

// 修改用户信息
func updateUser(users []map[string]string) []map[string]string {
	var userId string
	fmt.Print("请输入要修改用户的ID号：")
	fmt.Scanln(&userId)
	if ok, index := isId(users, userId); ok == true{
		// 修改用户信息
		var isupdate string
		fmt.Printf("请确认是否修改ID为%v的用户(y/n)", userId)
		fmt.Scanln(&isupdate)
		if isupdate == "y"{
			var userName, userAddress, userTel string
			fmt.Print("请输入用户名：")
			fmt.Scanln(&userName)
			fmt.Print("请输入用户地址：")
			fmt.Scanln(&userAddress)
			fmt.Print("请输入用户手机号：")
			fmt.Scanln(&userTel)
			users[index]["name"] =  userName
			users[index]["address"] = userAddress
			users[index]["tel"] = userTel
			return users
		}else if isupdate == "n" {
			return users
		}
		fmt.Println("你输入的指令不正确")
		return users
	}
	fmt.Println("你输入的用户ID不存在！！！")
	return users
}