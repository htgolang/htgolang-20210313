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

//删除用户
func delUser() error {
	var id string
	id = input("请输入要删除的用户ID:")

	i, v := checkId(id)

	if v {
		fmt.Println(UserInfo[i])
		confirm := input("确认删除吗(y/n)？")
		if confirm == "y" {
			UserInfo = append(UserInfo[:i], UserInfo[i+1:]...)
			return nil
		} else if confirm == "n" {
			fmt.Println("取消删除")
			return nil
		} else {
			return errors.New("确认信息输入有误")
		}
	} else {
		return errors.New("用户不存在")
	}
}

func main() {
	if err := delUser(); err == nil {
		fmt.Println(UserInfo)
	} else {
		fmt.Printf("%v", err)
	}
}
