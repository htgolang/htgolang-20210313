package main

import "fmt"

func main() {
	left, right := 1, 2
	fmt.Println(left, right)
	tmp := left
	left = right
	right = tmp
	fmt.Println(left, right)

	left, right = right, left
	fmt.Println(left, right)
}
