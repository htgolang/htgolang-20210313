package command

import "strconv"

var Id int = 0
var Commands map[string]func() = map[string]func(){}

var Prompts [][2]string

func Register(f func(), desc string) {
	Id++
	id := strconv.Itoa(Id)
	Prompts = append(Prompts, [2]string{id, desc})
	Commands[id] = f
}
