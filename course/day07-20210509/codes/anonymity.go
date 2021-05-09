package main

import "fmt"

type GobPersistent struct{}

// 定义了Persistent 所有的接口，就叫GobPersistent实现了Persistent接口

func (p GobPersistent) Save(path string) error {
	fmt.Println("gob persistent save")
	return nil
}

func (p GobPersistent) Load(path string) error {
	fmt.Println("gob persistent load")
	return nil
}
func main() {
	var test = func() {

	}

	var user struct {
		ID   string
		Name string
	}

	var saver interface {
		Save(string) error
	} = GobPersistent{}

	fmt.Printf("%T, %#v\n", test, test)
	fmt.Printf("%T, %#v\n", user, user)
	fmt.Printf("%T, %#v\n", saver, saver)
}
