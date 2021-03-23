package main

import "fmt"

func main() {
	for n := 1; n < 10; n++ {
		for i := 1; i <= n; i++ {
			fmt.Printf("%d * %d = %d\t", i, n, n+i)
		}
		fmt.Println("")
	}

}
