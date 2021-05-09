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

type Dumper interface {
	Save([]User, string) error
	Dump()
}

func main() {
	var persistent Persistent = GobPersistent{}
	fmt.Printf("%#v\n", persistent)

	var saver Saver
	fmt.Printf("%T, %#v\n", saver, saver)
	// saver.Load("")

	saver = GobPersistent{}
	fmt.Printf("%T, %#v\n", saver, saver)
	saver.Save(nil, "")

	saver = persistent
	fmt.Printf("%T, %#v\n", saver, saver)

	var loader Loader = persistent
	fmt.Printf("%T, %#v\n", loader, loader)

	// var dumper Dumper = loader
	// fmt.Printf("%T, %#v\n", dumper, dumper)
}
