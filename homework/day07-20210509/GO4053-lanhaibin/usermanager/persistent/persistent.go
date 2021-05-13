package persistent

import "usermanager/model"

type UserPersistent interface {
	Encode(map[int]*model.User) error
	Decode(*map[int]*model.User) error
}

type TodoPersistent interface {
}
