package main

import "fmt"

func main() {
	user := map[string]string{}
	fmt.Println(user)
	//并非线程安全
}
