package main

import (
	"fmt"
	"strings"
)

func queryName(name string) {
	flag := false
	for i, v := range data {
		if strings.EqualFold(v["name"], name) {
			fmt.Println(data[i])
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
	for i, v := range data {
		if strings.EqualFold(v["addr"], addr) {
			fmt.Println(data[i])
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
	for i, v := range data {
		if strings.EqualFold(v["tel"], tel) {
			fmt.Println(data[i])
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
