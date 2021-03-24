package main

import "fmt"

func main() {
	// 打印长方形
	fmt.Println("长方形")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}

	// 打印正三角形
	fmt.Println("正三角")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, j*i)
		}
		fmt.Println()
	}

	// 打印倒三角形
	fmt.Println("倒三角")
	for i := 1; i <= 3; i++ {
		for j := i; j <= 3; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, j*i)
		}
		fmt.Println()
	}
}
