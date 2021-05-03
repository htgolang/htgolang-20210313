package service

import (
	"errors"
	"usermanager/model"
)

type UserService struct{}

func (UserService) Get(id int) (*model.User, error) {
	user, ok := model.Users[id]
	if !ok {
		return nil, errors.New("用户不存在!")
	}
	return user, nil
}

func (UserService) Add(user model.User) error {
	err := ValidUser(user)
	if err != nil {
		return err
	}
	id := model.UserId
	user.Id = id
	model.Users[id] = &user
	model.UserId++
	return nil
}

func (UserService) Delete(id int) error {
	var flag bool = false
	for _, user := range model.Users {
		if user.Id == id {
			delete(model.Users, id)
			flag = true
			break
		}
	}
	if flag {
		return nil
	}
	return errors.New("用户不存在!")
}

func (UserService) Query(s string) model.UserList {
	var qs model.UserList
	for _, user := range model.Users {
		if user.Contains(s) {
			qs = append(qs, user)
		}
	}
	return qs
}

func (UserService) Modify(id int, u model.User) error {
	user, ok := model.Users[id]
	if !ok {
		return errors.New("用户不存在!")
	}
	if len(u.Name) != 0 {
		user.Name = u.Name
	}
	if len(u.Addr) != 0 {
		user.Addr = u.Addr
	}
	if len(u.Tel) != 0 {
		user.Tel = u.Tel
	}
	return nil
}

func ValidUser(u model.User) error {
	if u.Name == "" {
		return errors.New("用户名不能为空!")
	}
	for _, user := range model.Users {
		if user.Name == u.Name && user.Id != u.Id {
			return errors.New("用户已存在!")
		}
	}
	return nil
}
