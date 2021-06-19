package user

import (
	"fmt"
	"strings"
)

var usersSlice = make([]map[string]string, 0)
var users = make(map[string]string)
var userName string
var address string
var phone string
var con string

func Add() {
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
}

func Delete() {
	var id string
	var del string
	fmt.Println("输入要输出的用户ID：")
	fmt.Scan(&id)
	for i, v := range usersSlice {
		if v["id"] == id {
			fmt.Printf("%v", v)
			fmt.Println("是否要删除？Y/N")
			fmt.Scan(&del)
			if con == "Y" {
				usersSlice[i] = make(map[string]string)
			} else {
				continue
			}
		} else {
			continue
		}
	}
}

func Modify() {
	var id string
	var del string
	fmt.Println("输入要修改的用户ID：")
	fmt.Scan(&id)
	for _, v := range usersSlice {
		if v["id"] == id {
			fmt.Printf("%v", v)
			fmt.Println("请输入修改用户名：")
			fmt.Scan(&userName)
			fmt.Println("请输入修改地址：")
			fmt.Scan(&address)
			fmt.Println("请输入修改电话号码：")
			fmt.Scan(&phone)
			fmt.Scan(&del)
			fmt.Println("是否要修改？Y/N")
			if con == "Y" {
				users["name"] = userName
				users["addr"] = address
				users["tel"] = phone
			} else {
				continue
			}
		} else {
			continue
		}
	}
}

func Query() {
	var subStr string
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
