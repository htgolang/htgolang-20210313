package auth

import (
	"errors"
	"usermanager/model"
)

func Authencation(username, password string) (*model.User, error) {
	for _, user := range model.Users {
		if user.Name == username && user.CheckPassword(password) {
			return user, nil
		}
	}
	return &model.User{}, errors.New("用户名或密码错误")
}
