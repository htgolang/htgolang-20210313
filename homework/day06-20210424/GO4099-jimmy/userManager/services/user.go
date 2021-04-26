package services

import (
	"fmt"
	"strconv"
	"strings"
	"userManager/model"
	"userManager/utils"
)

type UserService struct{}

func (UserService) Add() {
	users := decode()

	id := generateId(users)
	user := utils.InputUser(id, "add", "")
	users[id] = user

	encode(users)

	fmt.Println("用户添加成功")
	fmt.Println(user)
}

func (UserService) Modify() {
	users := decode()
	id, _ := strconv.Atoi(utils.Input("请输入要修改的用户id: "))
	if u, ok := users[id]; ok {
		fmt.Println(u)
		switch utils.Input("确认是否编辑y/yes: ") {
		case "y", "yes":
			users[id] = utils.InputUser(id, "", users[id].Password)
			encode(users)
			fmt.Println("成功修改用户信息")
			fmt.Println(users[id])
		default:
			fmt.Println("未执行修改动作")
		}
	} else {
		fmt.Println("用户未找到")
	}
}

func (UserService) Query() {
	users := decode()

	qc := utils.Input("请输入要查询的字符串: ")
	for id, user := range users {
		if strings.Contains(user.Name, qc) || strings.Contains(user.Addr, qc) || strings.Contains(user.Tel, qc) || strings.Contains(strconv.Itoa(id), qc) {
			fmt.Println(user)
		}
	}
	fmt.Println(strings.Repeat("*", 20))
}

func (UserService) Delete() {
	users := decode()
	id, _ := strconv.Atoi(utils.Input("请输入要删除的用户id: "))
	if u, ok := users[id]; ok {
		fmt.Println(u)
		switch utils.Input("确认是否删除y/yes: ") {
		case "y", "yes":
			delete(users, id)
			encode(users)
			fmt.Println("用户删除成功")
		default:
			fmt.Println("未执行删除动作")
		}
	} else {
		fmt.Println("用户未找到")
	}
}

func generateId(users map[int]model.User) int {
	maxId := 0
	for k := range users {
		if k > maxId {
			maxId = k
		}
	}
	return maxId + 1
}
