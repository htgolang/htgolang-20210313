package user

import (
	"fmt"
	"strings"
	"user/common"
)

func queryName(name string) {
	flag := false
	for i, v := range common.UserList {
		if strings.EqualFold(v["name"], name) {
			fmt.Println(common.UserList[i])
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println("没找到")
	}
}

func queryAddr(addr string) {
	flag := false
	for i, v := range common.UserList {
		if strings.EqualFold(v["addr"], addr) {
			fmt.Println(common.UserList[i])
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println("没找到")
	}
}

func queryTel(tel string) {
	flag := false
	for i, v := range common.UserList {
		if strings.EqualFold(v["tel"], tel) {
			fmt.Println(common.UserList[i])
			flag = true
			break
		}
	}
	if !flag {
		fmt.Println("没找到")
	}
}

func QueryUser() {
	var query string
	fmt.Println("请输入要查询的字段(name/addr/tel)：")
	fmt.Scanln(&query)
	switch query {
	case "name":
		fmt.Println("请输入name")
		name := ""
		fmt.Scanln(&name)
		queryName(name)
		break
	case "addr":
		fmt.Println("请输入addr")
		addr := ""
		fmt.Scanln(&addr)
		queryAddr(addr)
		break
	case "tel":
		fmt.Println("请输入tel")
		tel := ""
		fmt.Scanln(&tel)
		queryTel(tel)
		break
	}
}
