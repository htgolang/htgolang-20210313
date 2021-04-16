package user

import (
	"fmt"
	"strings"
	"user/common"
	"user/utils"
)

func handleLogin(name, passwd string) bool {
	for _, v := range common.UserList {
		if strings.EqualFold(name, v["Name"]) && strings.EqualFold(utils.Md5sum(passwd), v["PassWord"]) {
			fmt.Println("登录成功")
			return true
		}
	}
	return false
}

func Check() bool {
	userName := utils.StrInput("请输入用户名")
	passWord := utils.StrInput("请输入密码")
	if handleLogin(userName, passWord) {
		return true
	}
	return false
}
