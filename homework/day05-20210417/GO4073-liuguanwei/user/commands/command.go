package commands

//import "user/user"
import (
	"strconv"
)

var id int = 2

var Commands map[string]func() = map[string] func(){
	/*
	"2": user.Add,
	"3": user.Edit,
	"4": user.Del,
	"5": user.Query,
	*/
}

var Prompts [][2]string = [][2]string{
	/*
	{"2","add"},
	{"3","edit"},
	{"4","del"},
	{"5","query"},
	*/
}  

//Optimize
func Register(desc string, callback func()){
	Commands[strconv.Itoa(id)] = callback
	Prompts = append(Prompts, [2]string{strconv.Itoa(id),desc})
	id++
}