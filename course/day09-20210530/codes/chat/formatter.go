package main

import "fmt"

func main() {
	// "%5d" 5 => size
	size := 5
	fmt.Printf("%%(%d)d\n", size)
	fmt.Printf("%%%dd\n", size)

	fmt.Printf("%05d\n", 1)
	fmt.Printf("%5d\n", 1)
	fmt.Printf("%5d\n", 1000000)
}
