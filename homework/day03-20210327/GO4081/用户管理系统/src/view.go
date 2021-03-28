/**
* @Author: fengzhilaoling
* @Date: 2021/3/28 10:50
* 查看所有用户信息
**/
package main

import (
	"fmt"
)

// 查看所有的用户信息
func viewUser(users []map[string]string)  {
	for _, value := range users{
		fmt.Printf("用户的ID：%v  用户的姓名：%v  用户的地址：%v, 用户的手机号：%v\n", value["id"],
			value["name"], value["address"], value["tel"])
	}
}