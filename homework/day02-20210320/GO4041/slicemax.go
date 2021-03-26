package main

import "fmt"

func main() {
	nums := []int{9, 8, 19, 10, 2, 8}

	maxValue, maxvIndex := 0, 0
	for i, j := range nums {
		if j > maxValue {
			maxValue = j
			maxvIndex = i
		}
	}
	fmt.Println("切片最大值：", maxValue, "切片最大值索引", maxvIndex)

	maxValuev2, maxvIndexv2 := 0, 0
	for i, j := range nums {
		if j < maxValue && j > maxValuev2 {
			maxValuev2 = j
			maxvIndexv2 = i
		}
	}
	fmt.Println("切片第二大值：", maxValuev2, "切片第二大值索引", maxvIndexv2)

}
