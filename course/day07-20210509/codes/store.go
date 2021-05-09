package main

import "fmt"

type Persistent interface {
	Save(interface{}) error
	Load(string, interface{}) error
}

type PersistentV2 interface {
	Save([]interface{}) error
	Load(string, []interface{}) error
}

type User struct {
	ID string
}

func main() {
	var emptySlice []interface{}

	var empty interface{}
	fmt.Println(emptySlice, empty)

	var users = []User{{"1"}, {"2"}}

	empty = users
	fmt.Println(empty)

	// emptySlice = users // 切片赋值元素类型要一致
	// fmt.Println(emptySlice)
	emptySlice = []interface{}{users[0], users[1]}
	fmt.Println(emptySlice)

}
