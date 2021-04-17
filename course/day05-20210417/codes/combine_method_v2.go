package main

import "fmt"

type Address struct {
	Name string
}

func (address Address) Print() {
	fmt.Println(address.Name)
}

type User struct {
	Name string
	Address
}

func (user User) Print2() {
	fmt.Println(user.Name)
}

func main() {
	user := User{Name: "kk", Address: Address{Name: "陕西省"}}
	user.Print()         // 陕西 user.Address.Print() 接收者 为user.Address // 陕西省
	user.Address.Print() //陕西省

	user.Print2() // kk
}
