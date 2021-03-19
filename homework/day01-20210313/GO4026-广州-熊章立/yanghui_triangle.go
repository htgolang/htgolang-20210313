package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var sj [10][10]int
	for j := 0; j < a; j++ {
		sj[j][0] = 1
		sj[j][j] = 1
	}

	for j := 2; j < a; j++ {
		for n := 1; n < j; n++ {
			sj[j][n] = sj[j-1][n-1] + sj[j-1][n]
		}
	}

	for q := 0; q < a; q++ {
		fmt.Print(q)
		for m := a - q; m >= 0; m-- {
			fmt.Print(" ")
		}

		for n := 0; n <= q; n++ {
			fmt.Print(sj[q][n], " ")
		}
		fmt.Println()
	}
}
