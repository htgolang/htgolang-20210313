package main

import "fmt"

func main() {
	a1, a2 := 1, 2

	fmt.Println(a1, a2)
	a1, a2 = a2, a1
	fmt.Println(a1, a2)
}
