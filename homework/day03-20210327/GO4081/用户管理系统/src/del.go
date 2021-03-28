/**
* @Author: fengzhilaoling
* @Date: 2021/3/28 11:38
* 删除用户
**/
package main

import "fmt"

// 判断用户ID是否存在
func isId(users []map[string]string, id string) (bool, int) {
	for index, value := range users{
		if value["id"] == id{
			return true, index
		}
	}
	return false, -1
}

// 删除用户
func delUser(users []map[string]string) []map[string]string {
	var userId string
	fmt.Print("请输入要删除用户的ID号：")
	fmt.Scanln(&userId)
	if ok, index := isId(users, userId); ok == true{
		// 删除用户
		var isdel string
		fmt.Printf("请确认是否删除ID为%v的用户(y/n)", userId)
		fmt.Scanln(&isdel)
		if isdel == "y"{
			users = append(users[:index], users[index + 1:]...)
			return users
		}else if isdel == "n" {
			return users
		}
		fmt.Println("你输入的指令不正确")
		return users
	}
	fmt.Println("你输入的用户ID不存在！！！")
	return users
}