package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %-6d", j, i, j*i)
		}
		fmt.Println()
	}
}
