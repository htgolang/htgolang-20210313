package main

import "fmt"

func main() {

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v\t", j, i, i*j)

		}
		fmt.Printf("\n")
	}
}
