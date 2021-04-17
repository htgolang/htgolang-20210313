package main

import "fmt"

type Address struct {
	Province string
	City     string
	Street   string
}

type User struct {
	Id      int
	Name    string
	Tel     string
	Address // 匿名嵌入 自定义属性名Address(与类型名一致), 只能被匿名嵌入一次
}

// 属性名与匿名嵌入对象属性名相同
type User2 struct {
	Id     int
	Name   string
	Tel    string
	Street string
	Address
}

// 匿名嵌入两个对象，两个对象中有相同的属性名（Province）

type Company struct {
	Name     string
	Province string
}

type User3 struct {
	Id     int
	Name   string
	Tel    string
	Street string
	Address
	Company
}

func main() {
	// 1. 定义变量
	var u User
	fmt.Printf("1. %#v\n", u) // u.Address
	// 2. 赋值
	u = User{Id: 11, Address: Address{"陕西", "西安", "高新区"}}
	fmt.Printf("2. %#v\n", u)
	// 3. 属性访问和修改
	fmt.Println("3.", u.Address)
	fmt.Println("4.", u.Address.Street)
	u.Address.Street = "未央区"
	fmt.Printf("5. %#v\n", u)

	// 当访问匿名嵌入对象的属性时可以省略嵌入对象的属性名
	fmt.Println("6.", u.Street)
	u.Street = "西咸新区"
	fmt.Printf("7. %#v\n", u)

	var u2 User2

	// 优先查找当前对象，若无再查找匿名嵌入对象中属性名
	fmt.Println("8.", u2.Street)
	u2.Street = "xxxx"
	fmt.Printf("9. %#v\n", u2)

	u2.Address.Street = "yyy"
	fmt.Println("10.", u2.Address.Street)

	fmt.Printf("11. %#v\n", u2)

	// 多个匿名嵌入对象具有相同属性名
	var u3 User3

	fmt.Printf("12. %#v\n", u3)

	// fmt.Println(u3.Province) // 不允许 Address, Company 产生了二义性, go限制开发者使用全路径指名使用哪个属性
	fmt.Println("13.", u3.Address.Province)
	fmt.Println("14.", u3.Company.Province)
}
