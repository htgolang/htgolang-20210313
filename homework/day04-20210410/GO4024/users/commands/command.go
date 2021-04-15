package commands

import (
	"strconv"
	_ "fmt"
	"users/user"
)



var Commands = map[string]func(){
	"1" : user.Add,
	"2" : user.Modify,
	"3" : user.Delete,
	"4" : user.Query,
}



var Prompts [][2]string =[][2]string {
	{"1","添加用户"},
	{"2","编辑用户"},
	{"3","删除用户"},
	{"4","查询用户"},
}

// func Prompt() {
// 	fmt.Println(`
// 欢迎使用xxx管理系统
// 0，退出
// 1, 添加用户
// 2，编辑用户
// 3，删除用户
// 4，查询用户

// `)}


//需要掌握，web中经常用到
//添加菜单项而不用修改别的代码，只是增加init
// 《一定要掌握》
var id int = 2
func Register(desc string, callback func()) {
	Commands[strconv.Itoa(id)] = callback
	Prompts = append(Prompts,[2]string{strconv.Itoa(id),desc})
}