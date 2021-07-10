package auth

import (
	"errors"
	"usermanager/model"
	"usermanager/service"
)

func Authencation(username, password string) (model.User, error) {
	user, err := service.Us.GetUserByName(username)
	if err == nil && user.CheckPassword(password) {
		return user, nil
	}
	return user, errors.New("用户名或密码错误")
}
