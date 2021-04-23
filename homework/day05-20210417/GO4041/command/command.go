package command

import "user/user"

// type us user.Users

var Commands map[string]func(*user.Users) = map[string]func(*user.Users){}

func Register(desc string, callback func(*user.Users)) {
	Commands[desc] = callback
}
