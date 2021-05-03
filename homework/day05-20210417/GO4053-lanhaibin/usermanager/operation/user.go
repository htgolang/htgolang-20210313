package operation

import (
	"fmt"
	"strings"
	"usermanager/common"
	"usermanager/service"
	"usermanager/utils"
)

type UserOperation struct{}

func (UserOperation) Add() {
	user := common.InputUser()
	us := service.UserService{}
	err := us.Add(user)
	if err != nil {
		fmt.Println(err)
	}
}

func (UserOperation) Delete() {
	id := utils.InputInt("请输入用户id:")
	us := service.UserService{}
	user, err := us.Get(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("请确认用户信息:")
	fmt.Println(user)
	ask := utils.Input("是否删除?(y/n)")
	if strings.ToLower(ask) != "y" && strings.ToLower(ask) != "yes" {
		return
	}
	err = us.Delete(id)
	if err != nil {
		fmt.Println(err)
	}
}

func (UserOperation) Query() {
	s := utils.Input("请输入查询字符串:")
	us := service.UserService{}
	ul := us.Query(s)
	if len(ul) != 0 {
		fmt.Println("查询到以下用户信息:")
		for _, user := range ul {
			fmt.Println(user)
		}
	} else {
		fmt.Println("未查询到任何用户!")
	}
}

func (UserOperation) Modify() {
	id := utils.InputInt("请输入要修改的用户id:")
	us := service.UserService{}
	user, err := us.Get(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)
	ask := utils.Input("是否修改?")
	if strings.ToLower(ask) != "y" && strings.ToLower(ask) != "yes" {
		return
	}
	u := common.InputUser()
	fmt.Println(u)
	ask = utils.Input("是否修改?")
	if strings.ToLower(ask) != "y" && strings.ToLower(ask) != "yes" {
		fmt.Println(ask)
		return
	}
	err = us.Modify(id, u)
	if err != nil {
		fmt.Println(err)
	}
}
