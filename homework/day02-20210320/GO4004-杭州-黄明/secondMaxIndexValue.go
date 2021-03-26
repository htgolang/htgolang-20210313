package main

import "fmt"

func main() {
	slice := []int{9, 8, 19, 10, 2, 8, 17}
	max_value := 0
	second_index, second_value := 0, 0
	for i, v := range slice {
		if v > max_value {
			max_value = v
		} else {
			if v < max_value && v > second_value {
				second_index = i
				second_value = v
			}
		}
	}
	fmt.Printf("最大索引值：%v, 第二大索引：%v, 第二大索引值：%v\n", max_value, second_index, second_value)
}
