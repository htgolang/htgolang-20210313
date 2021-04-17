package main

import "fmt"

type User struct {
	id    int
	name  string
	tel   string
	email string
}

// (值接收者)定义User的方法, 返回用户名
func (u User) GetName() string {
	// u: 调用者的值拷贝
	return "GetName:" + u.name
}

// (值接收者)修改名称的方法
func (u User) SetName(name string) {
	u.name = name
}

// 指针接收者
func (u *User) PsetName(name string) {
	u.name = name
}

func (User) No() {
	fmt.Println("no")
}

func main() {
	u := User{name: "kk"}

	fmt.Println(u.GetName()) //调用方法 返回值 = 对象.方法名称(参数)

	u.SetName("arun")

	fmt.Println(u.GetName()) // GetName: "kk" ? GetName: arun

	// 调用PsetName (指针)

	// 语法糖 接收者 指针接收者 调用 值对象 => 自动进行取引用
	u.PsetName("arune") // => (&u).PsetName("arun")

	fmt.Println(u.GetName()) // GetName: arun

	p := &User{name: "kk"}

	// GetName
	// 语法糖 接收者 值接收者调用 指针对象 => 自动进行解引用
	fmt.Println(p.GetName()) // (*p).GetName()

	// 语法糖在编译过程中进行转换

	p.PsetName("arune")

	fmt.Println(p.GetName())

	p.No()

}
