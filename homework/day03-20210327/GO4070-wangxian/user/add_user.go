package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var UserInfo = []map[string]string{
	{"id": "1", "name": "jack", "address": "北京", "tel": "13582344859"},
	{"id": "7", "name": "eric", "address": "河北", "tel": "18282343859"},
	{"id": "9", "name": "周瑜", "address": "上海", "tel": "15782344859"},
}

//生成用户ID
func generateID() string {
	id := 0
	for _, user := range UserInfo {
		i, _ := strconv.Atoi(user["id"])
		if i > id {
			id = i
		}
	}
	return strconv.Itoa(id)
}

//获取用户输入信息
func input(msg string) (value string) {
	fmt.Print(msg)
	fmt.Scanln(&value)
	return strings.TrimSpace(value)
}

//判断用户是否存在
func checkUser(name string) bool {
	for _, user := range UserInfo {
		if name == user["name"] {
			return true
		}
	}
	return false
}

//添加用户
func addUser() error {
	var name, addr, tel string
	name = input("请输入用户名:")
	addr = input("请输入地址:")
	tel = input("请输入电话:")

	if checkUser(name) {
		return errors.New("用户已存在")
	}

	UserInfo = append(UserInfo, map[string]string{"id": generateID(),
		"name":    name,
		"tel":     tel,
		"address": addr})
	return nil
}

func main() {
	if err := addUser(); err == nil {
		fmt.Println(UserInfo)
	} else {
		fmt.Printf("添加失败，%v", err)
	}

}
