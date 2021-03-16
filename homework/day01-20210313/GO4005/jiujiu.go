package main

import (
	"fmt"
)

func main() {
	var i,j,sum int
	for i = 1; i <= 9; i++ {
		for j = 1; j <= i; j++ {
			sum = i * j
			fmt.Printf("%d * %d = %-4d", i , j , sum)
		}
		fmt.Printf("\n")
	}
}