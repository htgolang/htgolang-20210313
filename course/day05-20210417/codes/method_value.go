package main

import "fmt"

type User struct {
	Name string
}

func (u User) GetName() string {
	return u.Name
}

func (u *User) PGetName() string {
	return u.Name
}

func main() {
	u := User{"xxxx"}
	method1 := u.GetName // u -> 值接收者.GetName
	fmt.Printf("%T\n", method1)
	fmt.Println(method1()) // xxxxx
	u.Name = "yyyyy"
	fmt.Println(method1()) // xxxx ? yyyy

	method2 := u.PGetName  // &u => 指针接收者.GetName
	fmt.Println(method2()) // yyy

	u.Name = "zzzz"
	fmt.Println(method2())

}
