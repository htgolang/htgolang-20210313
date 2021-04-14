package user

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

func Add() error {
	name := input("请输入您的姓名:")
	sex := input("请输入您的性别:")
	age := input("请输入您的年龄:")
	addr := input("请输入您的地址:")
	tel := input("请输入您的电话:")
	// if name == "jone" || name == "james" || name == "lucy" {
	// 	return fmt.Errorf("用户名%s已存在", name)
	// }

	if name == "jone" || name == "james" || name == "lucy" {
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

	if err := Add(); err == nil {
		fmt.Println(users)
	} else {
		fmt.Println("添加失败", err)
	}
	fmt.Println(users)
	return nil
}

func Del() {
	var id string
	fmt.Printf("请输入要删除的用户Id: ")
	fmt.Scan(&id)

	for userIndex, user := range users {
		value, ok := user["id"]
		if ok && value == id {
			fmt.Println(user)
			var isDel string
			fmt.Printf("是否删除 y/n: ")
			fmt.Scan(&isDel)
			if isDel == "y" {
				users = append(users[:userIndex], users[userIndex+1:]...)
				fmt.Println(users)
			}
		}
	}
}

func Mod() {
	var id, name, sex, age, addr, tel string

	fmt.Printf("请输入要修改的用户Id: ")
	fmt.Scan(&id)

	for _, user := range users {
		value, ok := user["id"]
		if ok && value == id {
			fmt.Println(user)
			var isModify string
			fmt.Printf("是否确认修改 y/n: ")
			fmt.Scan(&isModify)
			if isModify == "y" {
				fmt.Println("请输入新的用户名:")
				fmt.Scan(&name)
				fmt.Println("请输入新的性别:")
				fmt.Scan(&sex)
				fmt.Println("请输入新的年龄:")
				fmt.Scan(&age)
				fmt.Println("请输入新的地址:")
				fmt.Scan(&addr)
				fmt.Println("请输入新的联系电话:")
				fmt.Scan(&tel)
				user["name"] = name
				user["sex"] = sex
				user["age"] = age
				user["addr"] = addr
				user["tel"] = tel
			}

			fmt.Println(users)
		}
	}
}

func Query() {
	var queryStr string
	fmt.Printf("请输入要查询的ID: ")
	fmt.Scan(&queryStr)
	for _, user := range users {
		for _, userInfo := range user {
			if userInfo == queryStr {
				fmt.Printf("%s %s %s %s %s %s  \n", user["id"], user["name"], user["sex"], user["age"], user["addr"], user["tel"])
			}
		}
	}

}
