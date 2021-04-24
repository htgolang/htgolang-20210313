package services

import (
	"fmt"
	"mgr/models"
	"mgr/utils"
)

var users = []*models.User{
	{1, "arun", "北京", "xxx1231213"},
	{2, "lujianfeng", "北京", "152xxxxxxxx"},
}

func genId() int {
	id := 0
	for _, user := range users {
		if user.ID > id {
			id = user.ID
		}
	}
	return id + 1
}

type UserService struct{}

func (UserService) Add() {
	name := utils.Input("请输入用户名:")
	addr := utils.Input("请输入地址:")
	tel := utils.Input("请输入联系方式:")

	// 验证 用户名不能重复
	if name == "kk" {
		fmt.Printf("用户名%s已存在", name)
		return
	}

	users = append(users, &models.User{
		genId(),
		name,
		addr,
		tel,
	})
}

func (UserService) Modify() {
	fmt.Println("modify")
}

func (UserService) Delete() {
	fmt.Println("delete")
}

func (UserService) Query() {
	fmt.Println("query")
	fmt.Println(users)
}
