package services

import (
	"cmdb/modules"
	"cmdb/utils"
	"fmt"
	"strconv"
	"strings"
)

var users []modules.User = []modules.User{
	{Id: 0, Name: "kk", Addr: "西安", Tel: "12325465", Passwd: "123456"},
	{Id: 2, Name: "kk1", Addr: "西安", Tel: "12897645", Passwd: "123456"},
	{Id: 3, Name: "kk2", Addr: "西安", Tel: "168698890", Passwd: "123456"},
	{Id: 4, Name: "kk3", Addr: "西安", Tel: "19998989", Passwd: "123456"},
	{Id: 5, Name: "kk4", Addr: "西安", Tel: "1567789", Passwd: "123456"},
}

func Query(keyWord string) []modules.User {
	tmpUsers := []modules.User{}

	for _, v := range users {
		if strings.Contains(strconv.Itoa(v.Id), keyWord) || strings.Contains(v.Name, keyWord) || strings.Contains(v.Addr, keyWord) || strings.Contains(v.Tel, keyWord) {
			tmpUsers = append(tmpUsers, v)
		}
	}
	return tmpUsers
}

func QueryByID(uid int) []modules.User {
	tmpUser := []modules.User{}

	for _, v := range users {
		if v.Id == uid {
			tmpUser = append(tmpUser, v)
			break
		}
	}
	return tmpUser
}

//添加用户
//返回一个string，如果为空表示添加成功， 如果不为空，表示错误信息
func Add(user modules.User) string {
	index := utils.CheckUserExist(user.Id, users)
	if index != -1 {
		return "用户ID已存在"
	}

	if utils.CheckUserName(user.Name, users) {
		return "用户名已存在"
	}

	if utils.CheckTel(user.Tel, users) {
		return "号码已存在"
	}

	users = append(users, user)
	fmt.Println(users)
	return ""
}

func DelUser(uid int) string {
	index := utils.CheckUserExist(uid, users)
	if index == -1 {
		return "指定ID用户不存在"
	}

	users = append(users[:index], users[index+1:]...)
	fmt.Println(users)
	return ""
}

func Edit(user modules.User) string {
	index := utils.CheckUserExist(user.Id, users)
	if index == -1 {
		return "用户不存在"
	}

	//对username和tel检查的时候，把要编辑的用户排除在外
	allUser := append(users[:index], users[index+1:]...)

	if utils.CheckUserName(user.Name, allUser) {
		return "用户名已存在"
	}

	if utils.CheckTel(user.Tel, allUser) {
		return "号码已存在"
	}

	users[index] = user
	fmt.Println(users)
	return ""
}
