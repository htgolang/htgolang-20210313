package user

import (
	"fmt"
	"strconv"
	"strings"
	"userManager/common"
)

var user map[int]map[string]string = map[int]map[string]string{}

func Add() {
	id := generateId()
	user[id] = common.InputUser()
	fmt.Println("添加用户成功")
	fmt.Printf("用户id:%d, 名字:%v, 地址:%v, 电话:%v\n", id, user[id]["name"], user[id]["addr"], user[id]["tel"])
}

func Delete() {
	id, _ := strconv.Atoi(common.Input("请输入要删除的用户id: "))
	if u, ok := user[id]; ok {
		fmt.Printf("用户id:%d, 名字:%v, 地址:%v, 电话:%v\n", id, u["name"], u["addr"], u["tel"])
		switch common.Input("确认是否删除y/yes: ") {
		case "y", "yes":
			delete(user, id)
			fmt.Println("用户删除成功")
		default:
			fmt.Println("未执行删除动作")
		}
	} else {
		fmt.Println("用户未找到")
	}
}

func Query() {
	qc := common.Input("请输入要查询的字符串: ")
	for id, u := range user {
		if strings.Contains(u["name"], qc) || strings.Contains(u["addr"], qc) || strings.Contains(u["tel"], qc) || strings.Contains(strconv.Itoa(id), qc) {
			fmt.Printf("用户id:%d, 名字:%v, 地址:%v, 电话:%v\n", id, u["name"], u["addr"], u["tel"])
		}
	}
	fmt.Println(strings.Repeat("*", 20))
}

func Edit() {
	id, _ := strconv.Atoi(common.Input("请输入要修改的用户id: "))
	if u, ok := user[id]; ok {
		fmt.Printf("用户id:%d, 名字:%v, 地址:%v, 电话:%v\n", id, u["name"], u["addr"], u["tel"])
		switch common.Input("确认是否编辑y/yes: ") {
		case "y", "yes":
			user[id] = common.InputUser()
			fmt.Println("成功修改用户信息")
			fmt.Printf("用户id:%d, 名字:%v, 地址:%v, 电话:%v\n", id, user[id]["name"], user[id]["addr"], user[id]["tel"])
		default:
			fmt.Println("未执行修改动作")
		}
	} else {
		fmt.Println("用户未找到")
	}
}

func generateId() int {
	maxId := 0
	for k := range user {
		if k > maxId {
			maxId = k
		}
	}
	return maxId + 1
}
