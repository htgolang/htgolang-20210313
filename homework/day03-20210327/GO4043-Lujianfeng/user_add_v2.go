package main

import (
	"fmt"
	"os"
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

	if err := addUser(); err == nil {
		fmt.Println(users)
	} else {
		fmt.Println("添加失败", err)
	}
	fmt.Println(users)
	return nil
}

func delUser() {
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

func modifyUser() {
	var id string
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
	//这块不太会写了，请教老师，谢谢！	
	// 这里判断id是否在，如果不在就说明不存在，如果在就调用添加函数继续添加
				addUser[name]=name ,
				addUser[sex]=sex ,
				addUser[age]=age,
				addUser[addr]=addr,
				addUser[tel]=tel ,
				}
			
				fmt.Println(users)
			}
		}
	}
}

func queryUser() {
	var queryStr string
	fmt.Printf("请输入要查询的ID: ")
	fmt.Scan(&queryStr)
	for _, user := range users {
		for userKey, userInfo := range user {
			if userInfo == queryStr {
				fmt.Printf("Id: %s, %s: %s \n", user["id"], userKey, userInfo)
			}
		}
	}

}

func main() {

	fmt.Printf(`
	1. 添加用户 ==> add
	2. 删除用户 ==> del
	3. 修改用户 ==> modify
	4. 查询用户 ==> query
	5. 退出 ==> quit
	`)

	for {
		var choiceOne string
		fmt.Printf("请选择你要做的操作: ")
		fmt.Scan(&choiceOne)
		switch choiceOne {
		case "add":
			addUser()
		case "del":
			delUser()
		case "modify":
			modifyUser()
		case "query":
			queryUser()
		case "quit":
			os.Exit(0)
		}
	}

}
