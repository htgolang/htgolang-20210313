package main

import "fmt"

func main() {
	//define the left of multiplication
	for i := 1; i < 10; i++ {
		//define the right of multiplication
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
			//To next line if i equals j
			if i == j {
				fmt.Println()
			}
		}
	}
}
