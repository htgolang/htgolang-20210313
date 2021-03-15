package main

import "fmt"

func main() {

	// 使用 for 循环打印正三角 乘法口诀表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, (i * j))
		}
		fmt.Println()
	}
}
