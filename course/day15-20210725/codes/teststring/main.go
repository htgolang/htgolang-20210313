package main

import "fmt"

type User struct {
	Id       int
	Name     string
	Password string
}

func (u User) String() string {
	return fmt.Sprintf("User{id=%d, name=%s}", u.Id, u.Name)
}

func main() {
	user := User{1, "kk", "123"}
	fmt.Println(user)
}
