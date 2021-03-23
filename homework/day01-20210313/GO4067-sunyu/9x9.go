package main

import "fmt"

func main() {
	var (
		i int
		j int = 1
	)
	for j = 1; j < 10; j++ {
		for i = 1; i <= j; i++ {
			var sum int
			sum = i * j
			fmt.Printf("%d * %d = %d ", i, j, sum)
		}
		fmt.Println("\n")
	}
}
