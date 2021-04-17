package main

import "fmt"

// 地址结构体
type Address struct {
	Province string
	City     string
	Street   string
}

// 定义用户信息
type User struct {
	Id    int
	Name  string
	Tel   string
	Email string
	Addr  Address // 嵌入Address结构体, 即定义Address结构体的属性, 组合Address结构体(命名嵌入)
	PAddr *Address
}

func main() {
	// 1. 定义
	var u User

	fmt.Printf("%#v\n", u) // u.Addr u.PAddr
	// 2. 赋值
	// var addr Address = Address{Province: "陕西", City: "西安", Street: "高新路"}
	// u = User{Id: 1, Name: "kk", Tel: "152xxxx", Email: "xxx@xxx.com", Addr: addr}

	u = User{Id: 1, Name: "kk", Tel: "152xxxx", Email: "xxx@xxx.com", Addr: Address{"陕西", "西安", "高新区"}, PAddr: &Address{"深圳", "深圳", "福田区"}}

	fmt.Printf("%#v\n", u) // u.Addr

	// 3. 属性访问和修改
	fmt.Println(u.Id)
	fmt.Println(u.Addr)
	fmt.Printf("%T, %#v\n", u.PAddr, u.PAddr)
	u.Id = 10
	u.Addr = Address{"北京", "北京", "海淀区"}
	u.PAddr = &Address{"北京", "北京", "昌平区"}
	fmt.Printf("%#v\n", u)
	fmt.Printf("%#v\n", u.PAddr)

	fmt.Println(u.Addr.Street)
	fmt.Println(u.PAddr.Street)
	u.Addr.Street = "朝阳区"
	u.PAddr.Street = "东城区"

	fmt.Printf("%#v\n", u)       // u.Addr
	fmt.Printf("%#v\n", u.PAddr) // u.Addr

}
