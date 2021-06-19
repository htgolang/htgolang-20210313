package service

import (
	"errors"
	"fmt"
	"usermanager/model"
)

type UserService struct{}

func (u UserService) Get(id int) (model.User, error) {
	if user, ok := model.Users[id]; ok {
		return *user, nil
	}
	return model.User{}, fmt.Errorf("user id %d not exists", id)
}

func (u UserService) Create(name, addr, tel, password string) error {
	if len(name) == 0 || len(password) == 0 {
		return errors.New("用户名和密码不能为空")
	}
	for _, user := range model.Users {
		if user.Name == name {
			return errors.New("用户已存在")
		}
	}
	user := model.User{
		Id:   model.GetUserId(),
		Name: name,
		Addr: addr,
		Tel:  tel,
	}
	user.SetPassword(password)
	model.Users[user.Id] = &user

	return nil
}

func (u UserService) Delete(id int) error {
	delete(model.Users, id)
	return nil
}

func (u UserService) Query(s string) (qs []model.User) {
	for _, user := range model.Users {
		if user.Contains(s) {
			qs = append(qs, *user)
		}
	}
	return
}

func (u UserService) Modify(id int, name, addr, tel string) error {
	user, ok := model.Users[id]
	if !ok {
		return fmt.Errorf("user id %d not exists", id)
	}
	if len(name) > 0 {
		for _, user := range model.Users {
			if user.Id != id && user.Name == name {
				return errors.New("用户名已存在")
			}
		}
		user.Name = name
	}
	if len(addr) > 0 {
		user.Addr = addr
	}
	if len(tel) > 0 {
		user.Tel = tel
	}
	return nil
}

// func (u UserService) ChangePassword(newpassword string) error {

// }

var Us UserService
