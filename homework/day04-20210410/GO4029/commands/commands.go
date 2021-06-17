package commands

import "strconv"

var id int = 2
var Command map[string]func() = map[string]func(){}
var Prompt [][2]string

func Register(desc string, callback func()) {
	Command[strconv.Itoa(id)] = callback
	Prompt = append(Prompt, [2]string{strconv.Itoa(id), desc})
	id++
}
