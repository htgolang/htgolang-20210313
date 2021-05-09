package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int    `kk:"yyy"`
	Name string `kk:"xxxx"`
}

func (u User) GetId() int {
	return u.Id
}

func (u *User) GetName() string {
	return u.Name
}

type Persistent interface {
	Save()
}

type DbPersistent struct{}

func (p DbPersistent) Save() {

}

func main() {
	var u = new(User)
	t := reflect.TypeOf(u)
	fmt.Printf("%T, %#v\n", t, t)
	fmt.Println(t.Name())
	fmt.Println(t.PkgPath())
	fmt.Println(t.Kind())
	fmt.Println(t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println(i, method.Name)
	}

	m, ok := t.MethodByName("GetName")
	fmt.Println(m, ok)

	m, ok = t.MethodByName("GetName1")
	fmt.Println(m, ok)

	st := t.Elem()

	fmt.Println(st.NumField())
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		fmt.Println(i, field.Name, field.Type.Kind(), field.Tag.Get("kk"))
	}

}
