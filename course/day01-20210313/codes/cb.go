package main

import "fmt"

func main() {
	fmt.Println("continue:")
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue // 跳过本次循环，下次继续
		}
		fmt.Println(i)
	}

	fmt.Println("break:")
	for i := 0; i < 5; i++ {
		if i == 3 {
			break // 结束循环
		}
		fmt.Println(i)
	}
}
