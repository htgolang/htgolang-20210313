package main

import (
	"fmt"
	"strconv"
)

var users = []map[string]string{
	{
		"id":   "1",
		"name": "jone",
		"sex":  "M",
		"age":  "32",
		"addr": "北京",
		"tel":  "xxx1231213",
	},
	{
		"id":   "2",
		"name": "james",
		"sex":  "M",
		"age":  "22",
		"addr": "上海",
		"tel":  "152xxxxxxxx",
	},
	{
		"id":   "3",
		"name": "lucy",
		"sex":  "W",
		"age":  "25",
		"addr": "深圳",
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

func addUser() error {
	name := input("请输入您的姓名:")
	sex := input("请输入您的性别:")
	age := input("请输入您的年龄:")
	addr := input("请输入您的地址:")
	tel := input("请输入您的电话:")
	// if name == "jone" || name == "james" || name == "lucy" {
	// 	return fmt.Errorf("用户名%s已存在", name)
	// }

	if name == "kk" {
		return fmt.Errorf("用户名%s已存在", name)
	}

	users = append(users, map[string]string{

		"id":   strconv.Itoa(genId()),
		"name": name,
		"sex":  sex,
		"age":  age,
		"addr": addr,
		"tel":  tel,
	})

	return nil
}
func main() {
	if err := addUser(); err == nil {
		fmt.Println(users)
	} else {
		fmt.Println("添加失败", err)
	}

}
