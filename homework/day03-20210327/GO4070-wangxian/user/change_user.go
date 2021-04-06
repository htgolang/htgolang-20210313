package main

import (
	"errors"
	"fmt"
	"strings"
)

var UserInfo = []map[string]string{
	{"id": "1", "name": "jack", "address": "北京", "tel": "13582344859"},
	{"id": "7", "name": "eric", "address": "河北", "tel": "18282343859"},
	{"id": "9", "name": "周瑜", "address": "上海", "tel": "15782344859"},
}

//获取用户输入信息
func input(msg string) (value string) {
	fmt.Print(msg)
	fmt.Scanln(&value)
	return strings.TrimSpace(value)
}

//判断id是否存在
func checkId(id string) (int, bool) {
	for i, user := range UserInfo {
		if id == user["id"] {
			return i, true
		}
	}
	return 0, false
}

//修改用户信息
func changUser() error {
	var n_name, n_addr, n_tel string
	var id string
	id = input("请输入要修改的用户id:")

	i, v := checkId(id)

	if v {
		n_name = input("请输入新的用户名:")
		n_addr = input("请输入新的地址:")
		n_tel = input("请输入新的电话:")

		confirm := input("是否确认修改(y/n)?")

		if confirm == "y" {
			UserInfo[i]["name"] = n_name
			UserInfo[i]["address"] = n_addr
			UserInfo[i]["tel"] = n_tel
			// fmt.Println(UserInfo)
			return nil
		} else if confirm == "n" {
			fmt.Println("取消修改")
			return nil
		} else {
			return errors.New("确认信息输入有误")
		}
	} else {
		return errors.New("id错误，用户不存在")
	}

}

func main() {
	if err := changUser(); err == nil {
		fmt.Println(UserInfo)
	} else {
		fmt.Printf("%v", err)
	}

}
