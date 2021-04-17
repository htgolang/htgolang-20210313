package main

import "fmt"

type Address struct {
	province string
	city     string
	street   string
}

func (address *Address) SetProvince(province string) {
	address.province = province
}

func (address *Address) GetProvince() string {
	return "Address:" + address.province
}
func (address *Address) Name() {
	fmt.Println("address")
}

type Company struct{}

func (Company) Name() {
	fmt.Println("company")
}

type User struct {
	id   int
	name string
	addr Address
	Address
	Company
}

func (u *User) GetProvince() string {
	return "user"
}

func main() {
	var u User
	u.addr.SetProvince("shannxi")
	fmt.Printf("%#v\n", u)

	// 匿名嵌入
	u.Address.SetProvince("北京")
	fmt.Printf("%#v\n", u)
	u.SetProvince("北京1")
	fmt.Printf("%#v\n", u)1

	// 当前对象存在GetProvince
	fmt.Println(u.GetProvince())
	fmt.Println(u.Address.GetProvince())

	// 匿名嵌入多个对象具有同名方法，需要用全路径访问
	// u.Name() => 具有二义性
	u.Address.Name()
	u.Company.Name()
}
