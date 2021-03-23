package main

import "fmt"

// 打印九九乘法表

func main() {
	// 法一
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
	// 法二
	var a int = 1
	for a <= 9 {
		var b int = 1
		for b <= a {
			fmt.Printf("%d * %d = %d\t", b, a, a*b)
			b++
		}
		fmt.Println()
	a++
	}
}
