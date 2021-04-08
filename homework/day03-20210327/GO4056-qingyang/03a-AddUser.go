package main

import (
	"fmt"
	"strconv"
)

func AddUser() {
	var (
		name, addr, tel string
	)
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入地址：")
	fmt.Scanln(&addr)
	fmt.Println("请输入电话：")
	fmt.Scanln(&tel)
	id++
	data = append(data, map[string]string{
		"id":   strconv.Itoa(id),
		"name": name,
		"addr": addr,
		"tel":  tel,
	})
	fmt.Println("添加用户成功")
}
