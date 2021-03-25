package main

import "fmt"

func main() {
	slice := []int{9, 8, 19, 10, 2, 8}
	num := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > num {
			num = slice[i]
		}
	}
	fmt.Println("最大数字是:", num)
}
