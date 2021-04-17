package main

import "fmt"

type User struct {
	id int
}

func (user *User) P() {
	fmt.Println("P")
}

func (user User) V() {
	fmt.Println("V")
}

func (user *User) Id() {
	// if user == nil {
	// 	return
	// }
	user.id += 1
}

func Pfunc(user *User) {
	fmt.Println("pfunc")
}

func Vfunc(user User) {
	fmt.Println("vfunc")
}

func main() {
	// 只对 接收者进行语法糖处理
	u := User{}
	p := new(User)

	u.P() // 语法糖 自定取引用
	p.V() // 语法糖 自动解引用

	// 针对函数无取引用和解引用语法
	Pfunc(&u)
	Vfunc(*p)

	var pp *User // 一定要初始化

	pp.P()
	pp.Id()
}
