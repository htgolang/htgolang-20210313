package main

import "fmt"

type Address struct {
	Province string
	City     string
	Street   string
}

type User struct {
	Id       int
	Name     string
	Tel      string
	*Address // 匿名嵌入 自定义属性名Address(与类型名一致), 只能被匿名嵌入一次
}

func main() {
	// 1. 定义变量

	var u User
	fmt.Printf("1. %#v\n", u) // u.Address

	// 2. 赋值

	u = User{Id: 11, Address: &Address{"陕西", "西安", "高新区"}}
	fmt.Printf("2. %#v\n", u)         // u.Address
	fmt.Printf("3. %#v\n", u.Address) // u.Address

	// 3. 属性访问和修改
	fmt.Println("4.", u.Address.Street)
	u.Address.Street = "未央区"
	fmt.Println("5.", u.Address.Street)

	u.Street = "xxx"
	fmt.Println("6.", u.Street, u.Address.Street) // xxx

	u2 := u
	u2.Address.Street = "yyyy"
	fmt.Println("7.", u.Address) // xxx, yyyy?
}
