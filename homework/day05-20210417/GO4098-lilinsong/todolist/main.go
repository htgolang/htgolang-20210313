package main

import (
	"todolist/commands"
	_ "todolist/init"
)

func main() {
	go commands.BakFile()
	commands.Run()
}
