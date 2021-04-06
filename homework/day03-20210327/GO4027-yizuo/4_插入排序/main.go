package main

import "fmt"

// 插入排序
func insertSort(data []int) {
	/*
		[8 9 19 10 2 8]
		[8 9 10 19 2 8]
		[2 9 10 19 8 8]
		[2 8 10 19 9 8]
		[2 8 9 19 10 8]
		[2 8 9 10 19 8]
		[2 8 8 10 19 9]
		[2 8 8 9 19 10]
		[2 8 8 9 10 19]

		实际流程如上，插入排序，从左到右将现有位置的值与已在列表中的每个值进行比较，然后交换。
	*/
	for k, _ := range data {
		for i := 0; i < k; i++ {
			if data[k] < data[i] {
				data[k], data[i] = data[i], data[k]
			}
		}
	}
	fmt.Println(data)
}

func main() {
	data := []int{9, 8, 19, 10, 2, 8}
	insertSort(data)
}
