package main

import (
	"fmt"
)

func main() {
	// 打印9*9乘法口诀表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
			if j == i {
				fmt.Println()
			}
		}
	}
}
