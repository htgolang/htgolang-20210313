package main

import "fmt"

func main() {
	var i, j int
	for i = 0; i < 9; {
		i++
		for j = 0; j < i; {
			j++
			fmt.Printf("%d * %d => %v\t", i, j, i*j)
		}
		fmt.Println("")
	}

}
