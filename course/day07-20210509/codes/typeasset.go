package main

import "fmt"

type User struct{}

type Persistent interface {
	Save([]User, string) error
	Load(string) ([]User, error)
}

type GobPersistent struct {
	Version string
}

func (p GobPersistent) Save(users []User, path string) error {
	fmt.Println("save")
	return nil
}

func (p GobPersistent) Load(path string) ([]User, error) {
	fmt.Println("load")
	return nil, nil
}

func (p GobPersistent) Test() {
	fmt.Println("test")
}

type CsvPersistent struct {
	Version string
}

func (p CsvPersistent) Save(users []User, path string) error {
	fmt.Println("save")
	return nil
}

func (p CsvPersistent) Load(path string) ([]User, error) {
	fmt.Println("load")
	return nil, nil
}

type Saver interface {
	Save([]User, string) error
}

func main() {
	var persistent Persistent = GobPersistent{"1.1.1"}

	// 断言 => 接口
	// 对应类型的变量, ok:= 接口类型变量.(类型)

	gobPersistent, ok := persistent.(GobPersistent)
	fmt.Printf("%T, %#v, %T, %#v\n", ok, ok, gobPersistent, gobPersistent)
	gobPersistent.Test()
	fmt.Println(gobPersistent.Version)

	// user, ok := persistent.(User)
	// fmt.Println(user, ok)
	csvPersistent, ok := persistent.(CsvPersistent)
	fmt.Printf("%T, %#v, %T, %#v\n", ok, ok, csvPersistent, csvPersistent)

	var saver Saver = persistent

	// Persistent, GobPersistent, CsvPersistent
	p, ok := saver.(Persistent)
	fmt.Printf("%T, %#v, %T, %#v\n", ok, ok, p, p)
	p.Load("")

	g, ok := saver.(GobPersistent)
	fmt.Printf("%T, %#v, %T, %#v\n", ok, ok, g, g)
	g.Test()
	fmt.Println(g.Version)

	c, ok := saver.(CsvPersistent)
	fmt.Printf("%T, %#v, %T, %#v\n", ok, ok, c, c)

}
