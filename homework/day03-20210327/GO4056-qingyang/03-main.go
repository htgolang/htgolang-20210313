package main

import "fmt"

var (
	id       int
	data     []map[string]string
	register map[string]func()
)

func init() {
	register = make(map[string]func())
	id = 0
	register["a"] = AddUser
	register["b"] = DeleteUser
	register["c"] = EditUser
	register["d"] = QueryUser
}

func welcome() {
	var c string
	fmt.Println("a.添加用户")
	fmt.Println("b.删除用户")
	fmt.Println("c.修改用户")
	fmt.Println("d.查询用户")
	fmt.Print("请输入操作序号：")
	fmt.Scanln(&c)
	register[c]()
}

func main() {
	for {
		welcome()
	}
}
