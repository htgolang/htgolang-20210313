package common

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"usermanager/model"
	"usermanager/utils"
)

func InputUser() (model.User, error) {
	user := model.User{}
	user.Name = utils.Input("请输入用户名:")
	user.Addr = utils.Input("请输入地址:")
	user.Tel = utils.Input("请输入电话号码:")
	return user, nil
}

func InputNewUser() (model.User, error) {
	user, _ := InputUser()
	password := utils.InputPassword("请输入密码:")
	cpassword := utils.InputPassword("请确认密码:")
	if password != cpassword {
		return user, errors.New("密码不一致!")
	}
	hash := sha512.New()
	hash.Write([]byte(password))
	user.Password = fmt.Sprintf("%x", hash.Sum([]byte{}))
	return user, nil
}
