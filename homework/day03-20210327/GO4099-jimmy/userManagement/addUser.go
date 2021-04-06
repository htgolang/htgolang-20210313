package main

import (
	"fmt"
	"strconv"
)

func addUser() {
	var name, addr, tel string
	fmt.Print("请输入用户名: ")
	fmt.Scan(&name)
	fmt.Print("请输入地址: ")
	fmt.Scan(&addr)
	fmt.Print("请输入电话: ")
	fmt.Scan(&tel)
	id, err := generateId()
	if err != nil {
		fmt.Println(err)
		return
	}
	users = append(users, map[string]string{
		"id": id,
		"name": name,
		"addr": addr,
		"tel": tel}, )
	fmt.Println(users)
	fmt.Println("添加用户成功")
}

func generateId() (string, error) {
	maxId := 0
	for _,v := range users {
		id, err := strconv.Atoi(v["id"])
		if err != nil {
			return  "",fmt.Errorf("%v转换为数字失败\n", v["id"])
		}
		if id > maxId {
			maxId = id
		}
	}
	return strconv.Itoa(maxId + 1), nil
}