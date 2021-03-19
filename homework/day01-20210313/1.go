package main

import "fmt"

func main() {
	//打印9*9
	var product int
	for a := 1; a < 10; a++ {
		for b := 1; b < 10; b++ {
			if a < b {
				break
			}
			product = a * b
			fmt.Printf("%d*%d=%d ", b, a, product)
		}
		fmt.Printf("\n")
	}
}
