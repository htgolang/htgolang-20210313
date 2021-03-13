package main

import "fmt"

func main() {
	// for expr {}

	var confirm string = "y"
	var i int = 0 // confirm == "y" i 自增1 并且打印
	//否则退出循环

	for confirm == "y" {
		fmt.Println(i)
		i++
		fmt.Print("继续吗?")
		fmt.Scan(&confirm)
	}

	fmt.Println("over", i)

}
