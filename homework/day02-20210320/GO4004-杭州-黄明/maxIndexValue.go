package main

import "fmt"

func main() {
	slice := []int{9, 8, 19, 10, 2, 8}
	max_index, max_value := 0, 0
	for i, v := range slice {
		if v > max_value {
			max_value = v
			max_index = i
		}
	}
	fmt.Printf("最大索引：%v，最大索引值：%v\n", max_index, max_value)
}
