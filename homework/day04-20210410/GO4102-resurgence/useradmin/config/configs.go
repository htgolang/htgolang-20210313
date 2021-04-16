package config

import (
	"strconv"
)

const PAS string = "202cb962ac59075b964b07152d234b70"

var ID = 1

type f func(*[]*User) error

type User struct {
	ID          int
	Name, Phone string
}

var FuncMap = map[string]f{}

var MenuArr = [][2]string{}

func Register(desc string, callback f) {
	sID := strconv.Itoa(ID)
	FuncMap[sID] = callback
	MenuArr = append(MenuArr, [2]string{sID, desc})
	ID++
}
