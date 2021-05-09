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

func main() {
	var persistent Persistent = GobPersistent{}
	// var persistent = GobPersistent{}
	fmt.Printf("%#v\n", persistent)

	persistent.Save(nil, "")
	persistent.Load("")
	// persistent.Test()               // 不能调用Test方法，Persistent接口无Test方法
	// fmt.Println(persistent.Version) // 不能获取Version属性
}
