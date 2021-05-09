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

type Storer struct {
	Persistent Persistent
}

type StorerV2 struct {
	Persistent
}

func main() {
	storer := Storer{GobPersistent{}}
	storer.Persistent.Load("")
	storer.Persistent.Save(nil, "")

	storerV2 := StorerV2{GobPersistent{}}
	storerV2.Persistent.Load("")
	storerV2.Persistent.Save(nil, "")

	storerV2.Load("")
	storerV2.Save(nil, "")
}
