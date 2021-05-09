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

type Saver interface {
	Save([]User, string) error
}

type Loader interface {
	Load(string) ([]User, error)
}

// 匿名嵌入
type PersistentV2 interface {
	Saver
	Loader
}

type PersistentV3 interface {
	Saver
	Loader
	Dump()
}

func main() {
	var persistent Persistent = GobPersistent{}
	fmt.Printf("%T, %#v\n", persistent, persistent)

	var persistentV2 PersistentV2 = GobPersistent{}
	fmt.Printf("%T, %#v\n", persistentV2, persistentV2)
	persistentV2.Load("")
	persistentV2.Save(nil, "")
}
