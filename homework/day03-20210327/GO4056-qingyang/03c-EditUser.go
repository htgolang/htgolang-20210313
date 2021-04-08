package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EditUser() {
	var (
		flag, name, addr, tel string
		id                    int
	)
	fmt.Println("请输入要删除的用户id：")
	fmt.Scanln(&id)
	for i, v := range data {
		val, _ := strconv.Atoi(v["id"])
		if val == id+1 {
			fmt.Println("用户存在，是否进行修改(y/yes/Y/YES)")
			fmt.Scanln(&flag)
			if strings.EqualFold(flag, "y") || strings.EqualFold(flag, "yes") || strings.EqualFold(flag, "Y") || strings.EqualFold(flag, "YES") {
				fmt.Println("请输入姓名：")
				fmt.Scanln(&name)
				fmt.Println("请输入地址：")
				fmt.Scanln(&addr)
				fmt.Println("请输入电话：")
				fmt.Scanln(&tel)
				data[i]["name"] = name
				data[i]["addr"] = addr
				data[i]["tel"] = tel
			} else {
				break
			}
		} else {
			fmt.Println("不存在该用户")
			break
		}
	}
}
