package main

import "fmt"

type User struct {
	Id   int
	Name string
}

// 定义接口Persistent
// 行为: Save保存User切片, Load加载User切片
type Persistent interface {
	Save([]User, string) error
	Load(string) ([]User, error)
}

type GobPersistent struct {
	GVersion string
}

// 定义了Persistent 所有的接口，就叫GobPersistent实现了Persistent接口

func (p GobPersistent) Save(users []User, path string) error {
	fmt.Println("gob persistent save")
	return nil
}

func (p GobPersistent) Load(path string) ([]User, error) {
	fmt.Println("gob persistent load")
	return nil, nil
}

type CsvPersistent struct {
	CVersion string
}

func (p CsvPersistent) Save(users []User, path string) error {
	fmt.Println("csv persistent save")
	return nil
}

func (p CsvPersistent) Load(path string) ([]User, error) {
	fmt.Println("csv persistent load")
	return nil, nil
}

func call(persistent Persistent) {
	fmt.Println("call:")
	persistent.Save(nil, "")
	persistent.Load("")
}

type JsonPersistent struct {
	Jversion string
}

func (p *JsonPersistent) Save(users []User, path string) error {
	fmt.Println("json persistent save")
	return nil
}

func (p *JsonPersistent) Load(path string) ([]User, error) {
	fmt.Println("json persistent load")
	return nil, nil
}

type ProtobufPersistent struct {
	Pversion string
}

func (p *ProtobufPersistent) Save(users []User, path string) error {
	fmt.Println("protobuf persistent save")
	return nil
}

func (p ProtobufPersistent) Load(path string) ([]User, error) {
	fmt.Println("protobuf persistent load")
	return nil, nil
}

type Other struct{}

func (p *Other) Save(users []User, path string) error {
	fmt.Println("other persistent save")
	return nil
}

func (p Other) Load(path string) ([]User, error) {
	fmt.Println("other persistent load")
	return nil, nil
}

func main() {
	var persistent Persistent = &ProtobufPersistent{"1.1.1"}

	if g, ok := persistent.(GobPersistent); ok {
		fmt.Println("g", g, ok)
	} else if c, ok := persistent.(CsvPersistent); ok {
		fmt.Println("c", c, ok)
	} else if j, ok := persistent.(*JsonPersistent); ok {
		fmt.Println("j", j, ok)
	} else if p, ok := persistent.(*ProtobufPersistent); ok {
		fmt.Println("p", p, ok)
	} else {
		fmt.Println("error")
	}

	/*
		switch 接口类型变量.(type) {
		case GobPersistent:
		case CsvPersistent:
		case *JsonPersistent:
		case *ProtobufPersistent:
		default:
		}

		switch p := 接口类型变量.(type) {
		case GobPersistent:
		case CsvPersistent:
		case *JsonPersistent:
		case *ProtobufPersistent:
		default:
		}
	*/

	switch persistent.(type) {
	case GobPersistent:
		fmt.Println("g")
	case CsvPersistent:
		fmt.Println("c")
	case *JsonPersistent:
		fmt.Println("j")
	case *ProtobufPersistent:
		fmt.Println("p")
	default:
		fmt.Println("err")
	}

	switch p := persistent.(type) {
	case GobPersistent:
		fmt.Println("g", p.GVersion)
	case CsvPersistent:
		fmt.Println("c", p.CVersion)
	case *JsonPersistent:
		fmt.Println("j", p.Jversion)
	case *ProtobufPersistent:
		fmt.Println("p", p.Pversion)
	default:
		fmt.Println("err")
	}
}
