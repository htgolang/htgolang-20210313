package main

import "fmt"

func main() {
	// 打印杨辉三角
	var counts int = 10

	for i := 0; i < counts; i++ {
		number := 1
		for k := 0; k < counts-i; k++ {
			fmt.Print(" ")
		}

		for j := 0; j <= i; j++ {
			fmt.Printf("%5d", number)
			number = number * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}
