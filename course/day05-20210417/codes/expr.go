package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name string
}

func (u User) GetName() string {
	return u.Name
}

/*
func (u *User) GetName() string {

}
*/

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) PsetName(name string) {
	u.Name = name
}

func main() {
	// 定义 => T 类型
	// 调用 => T(值/指针)变量/对象/实例
	// 访问方法
	// T.方法名
	// 定义对象 对象.方法名

	// 方法表达式 类型
	// 值类型接收者
	f := User.GetName
	fmt.Printf("%T\n", f)
	fmt.Println(f(User{Name: "kk"}))

	fmt.Println(User.GetName(User{Name: "xxxx"}))

	f2 := User.SetName
	fmt.Printf("%T\n", f2)
	f2(User{Name: "xxxx"}, "yyyy")

	//指针类型接收者
	f3 := (*User).PsetName //User.PsetName 报错了
	fmt.Printf("%T\n", f3)

	u := User{}
	f3(&u, "kk")
	fmt.Println(f(u))

	f4 := (*User).GetName // 如果定义了值接收者方法，GO自动生成一个对应的指针接收者方法

	fmt.Printf("%T\n", f4)

	fmt.Println(f4(&u))
	fmt.Println(strings.Repeat("*", 10))

	// 方法值
	u = User{Name: "kk"}
	m1 := u.GetName

	fmt.Printf("%T\n", m1)
	fmt.Println(m1())

	p := &User{Name: "kk"}
	m2 := p.GetName // 语法糖 解引用
	fmt.Printf("%T\n", m2)
	fmt.Println(m2())

	m3 := u.PsetName // 语法糖 取引用
	fmt.Printf("%T\n", m3)
	m3("xxxx")
	fmt.Println(u) // xxx ? kk

	m4 := p.PsetName
	fmt.Printf("%T\n", m4)
	m4("yyyy")
	fmt.Println(p)

}
