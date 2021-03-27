package main

import (
	"fmt"
	"strconv"
)

var users = []map[string]string{
	{
		"id":   "1",
		"name": "arun",
		"addr": "北京",
		"tel":  "xxx1231213",
	},
	{
		"id":   "2",
		"name": "lujianfeng",
		"addr": "北京",
		"tel":  "152xxxxxxxx",
	},
}

func genId() int {
	id := 0
	for _, user := range users {
		i, _ := strconv.Atoi(user["id"])
		if i > id {
			id = i
		}
	}
	return id + 1
}

func input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return text
}

func add() error {
	name := input("请输入用户名:")
	addr := input("请输入地址:")
	tel := input("请输入联系方式:")

	// 验证 用户名不能重复
	if name == "kk" {
		return fmt.Errorf("用户名%s已存在", name)
	}

	users = append(users, map[string]string{
		"id":   strconv.Itoa(genId()),
		"name": name,
		"addr": addr,
		"tel":  tel,
	})
	return nil
}

func main() {
	if err := add(); err == nil {
		fmt.Println(users)
	} else {
		fmt.Println("添加失败", err)
	}
}
