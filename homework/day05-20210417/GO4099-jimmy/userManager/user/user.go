package user

import (
	"fmt"
	"strconv"
	"strings"
	"userManager/common"
	"userManager/model"
)

var userList = model.UserList{Users: make(map[int]model.User)}

func Add() {
	id := generateId()
	user := common.InputUser(id)
	userList.Users[id] = *user
	fmt.Println("用户添加成功")
	fmt.Println(user)
}

func Query() {
	qc := common.Input("请输入要查询的字符串: ")
	for id, u := range userList.Users {
		if strings.Contains(u.Name, qc) || strings.Contains(u.Addr, qc) || strings.Contains(u.Tel, qc) || strings.Contains(strconv.Itoa(id), qc) {
			fmt.Println(u)
		}
	}
	fmt.Println(strings.Repeat("*", 20))
}

func Delete() {
	id, _ := strconv.Atoi(common.Input("请输入要删除的用户id: "))
	if u, ok := userList.Users[id]; ok {
		fmt.Println(u)
		switch common.Input("确认是否删除y/yes: ") {
		case "y", "yes":
			delete(userList.Users, id)
			fmt.Println("用户删除成功")
		default:
			fmt.Println("未执行删除动作")
		}
	} else {
		fmt.Println("用户未找到")
	}
}

func Modify() {
	id, _ := strconv.Atoi(common.Input("请输入要修改的用户id: "))
	if u, ok := userList.Users[id]; ok {
		fmt.Println(u)
		switch common.Input("确认是否编辑y/yes: ") {
		case "y", "yes":
			userList.Users[id] = *common.InputUser(id)
			fmt.Println("成功修改用户信息")
			fmt.Println(userList.Users[id])
		default:
			fmt.Println("未执行修改动作")
		}
	} else {
		fmt.Println("用户未找到")
	}
}

func generateId() int {
	maxId := 0
	for k := range userList.Users {
		if k > maxId {
			maxId = k
		}
	}
	return maxId + 1
}
