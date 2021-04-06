package main

import (
	"fmt"
	"strings"
)

var UserInfo = []map[string]string{
	{"id": "1", "name": "jack", "address": "北京", "tel": "13582344859"},
	{"id": "7", "name": "eric", "address": "河北", "tel": "18282343859"},
	{"id": "9", "name": "周瑜", "address": "上海", "tel": "15782344859"},
	{"id": "9", "name": "周强", "address": "上海", "tel": "13567921454"},
}

//获取用户输入信息
func input(msg string) (value string) {
	fmt.Print(msg)
	fmt.Scanln(&value)
	return strings.TrimSpace(value)
}

//查询用户信息
func searchUser() bool {
	flag := false
	keyWord := input("请输入要查询的关键词:")

	for _, user := range UserInfo {
		if strings.Contains(user["name"], keyWord) || strings.Contains(user["address"], keyWord) || strings.Contains(user["tel"], keyWord) {
			fmt.Printf("name:%s, address:%s, tel:%s\n", user["name"], user["address"], user["tel"])
			flag = true
		}
	}
	return flag
}

func main() {
	if !searchUser() {
		fmt.Println("没有相关用户")
	}

}
