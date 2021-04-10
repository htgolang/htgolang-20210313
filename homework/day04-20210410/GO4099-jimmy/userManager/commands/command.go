package commands

import (
	"strconv"
)

var id int = 2

var Commands map[string]func() = map[string]func(){}

var Prompts [][2]string = [][2]string{}

func Register(desc string, f func()) {
	Commands[strconv.Itoa(id)] = f
	Prompts = append(Prompts, [2]string{strconv.Itoa(id), desc})
	id++
}
