package main

import "fmt"

func main() {
	num := 6 // 0110
	// 右移动2 左移动2
	// 0001
	num >>= 2

	fmt.Println(num)
	// 0100
	num <<= 2
	fmt.Println(num)

	num = 6

	// 0001
	num >>= 2

	for i := 0; i < 2; i++ {
		num <<= 1 // 0111
		num |= 1
	}

	fmt.Println(num)
}
