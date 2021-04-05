package main

import (
	"fmt"
	"strings"
)

func main() {
	var usersSlice = make([]map[string]string, 0)
	var users = make(map[string]string)
	var userName string
	var address string
	var phone string
	var con string
	var subStr string
	for {
		fmt.Println("请输入用户名：")
		fmt.Scan(&userName)
		fmt.Println("请输入地址：")
		fmt.Scan(&address)
		fmt.Println("请输入电话号码：")
		fmt.Scan(&phone)
		userId := len(usersSlice) + 1
		users["id"] = fmt.Sprintf("%d", userId)
		users["name"] = userName
		users["addr"] = address
		users["tel"] = phone
		usersSlice = append(usersSlice, users)
		fmt.Printf("当前列表：%v", usersSlice)
		fmt.Println("是否继续添加，Y/N:")
		fmt.Scan(&con)
		if con == "Y" {
			continue
		} else {
			break
		}
	}
	fmt.Println("请查询字符串：")
	fmt.Scan(&subStr)
	for _, v := range usersSlice {
		name := v["name"]
		addr := v["addr"]
		tel := v["tel"]
		if strings.ContainsAny(name, subStr) || strings.Contains(addr, subStr) || strings.Contains(tel, subStr) {
			fmt.Printf("%v", v)
		} else {
			break
		}
	}
}
