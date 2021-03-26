package main

import "fmt"

func main() {
	array := [8]int{1, 4, 6, 5, 3, 10, 9, 8}
	//冒泡排序
	for i := 0; i < len(array); i++ {
		for j := 1; j < len(array)-i; j++ {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	fmt.Printf("%v\n", array)
	//打印第一大
	fmt.Printf("%v\n", array[len(array)-1])
	// 打印第二大
	fmt.Printf("%v\n", array[len(array)-2])
}
