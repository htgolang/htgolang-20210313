package main

import (
	"fmt"
	"strings"
)

func queryUser() {
	var text string
	fmt.Print("请输入要查找的任意属性字符串: ")
	fmt.Scan(&text)

	flag := false
	for _, v1 := range users {
		for _, v2 := range v1 {
			if strings.Contains(v2, text) {
				fmt.Println(v1)
				flag = true
				break
			}
		}
	}
	if !flag {
		fmt.Println("未找到有关此用户的信息")
	}
}