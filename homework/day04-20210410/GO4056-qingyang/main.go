package main

import (
	"fmt"
	_ "user/init"
	"user/user"
)

var (
	register map[string]func()
)

func init() {
	register = make(map[string]func())
	register["a"] = user.AddUser
	register["b"] = user.DeleteUser
	register["c"] = user.EditUser
	register["d"] = user.QueryUser
}

func welcome() {
	var c1 string
	str := `--------------------------------------------------
	***********用户管理系统***********
	    a    ）添加用户
  	    b    ）修改用户
	    c    ）删除用户
	    d    ）搜索信息
	********************************
	`
	fmt.Println(str)
	fmt.Print("请输入操作序号：")
	fmt.Scanln(&c1)
	register[c1]()
}

func main() {
	if !user.Check() {
		fmt.Println("身份错误")
		return
	}
	for {
		welcome()
	}
}
