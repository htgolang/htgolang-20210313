/**
* @Author: fengzhilaoling
* @Date: 2021/3/28 10:48
* 添加用户
**/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 设置要增加用户的ID的值
func setUserId(users []map[string]string) (string, error) {
	var nums []int
	for _, v := range users {
		num, error := strconv.Atoi(v["id"])
		if error != nil {
			return "0", error
		}
		nums = append(nums, num)
	}
	sort.Ints(nums)
	result := strconv.Itoa(nums[len(nums)-1] + 1)
	return result, nil
}

func addUser(users []map[string]string) []map[string]string {
	var userName, userAddress, userTel string
	fmt.Print("请输入用户名：")
	fmt.Scanln(&userName)
	fmt.Print("请输入用户所在地址：")
	fmt.Scanln(&userAddress)
	fmt.Print("请输入用户手机号：")
	fmt.Scanln(&userTel)
	userId, _ := setUserId(users)
	user := map[string]string{
		"id":      userId,
		"name":    userName,
		"address": userAddress,
		"tel":     userTel,
	}
	users = append(users, user)
	return users
}
