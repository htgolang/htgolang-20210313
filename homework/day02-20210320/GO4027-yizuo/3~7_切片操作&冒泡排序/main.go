package main

import (
	"fmt"
)

// 找出最大的值以及索引
func findLargestValue(slice []int) {
	var num = len(slice)
	var bubbleSortdLargestData = make([]int, num)
	copy(bubbleSortdLargestData, bubbleSort(slice))
	for k, v := range slice {
		if bubbleSortdLargestData[num-1] == v {
			fmt.Printf("切片中最大的值为：%v\n切片中最大的索引为: %v\n", v, k)
			break
		}
	}
}

// 找出第二大的值以及索引
func findSecodeValue(slice []int) {
	var num = len(slice)
	var bubbleSortdLargestData = make([]int, num)
	copy(bubbleSortdLargestData, bubbleSort(slice))
	for k, v := range slice {
		if bubbleSortdLargestData[num-2] == v {
			fmt.Printf("切片中第二大的值为：%v\n切片中第二大的索引为: %v\n", v, k)
			break
		}
	}
}

// 最大值与末尾值交换位置
func moveLargestValueToEnd(slice []int) {
	var num = len(slice)
	var sliceCopy = make([]int, num)
	var bubbleSortdLargestData = make([]int, num)
	copy(sliceCopy, slice)
	copy(bubbleSortdLargestData, bubbleSort(slice))
	for k, v := range slice {
		if bubbleSortdLargestData[num-1] == v {
			sliceCopy[k], sliceCopy[num-1] = sliceCopy[num-1], sliceCopy[k]
		}
	}
	fmt.Println("最大值与末尾值交换位置后切片列表：", sliceCopy)
}

// 倒数第二大的值与末尾值交换位置
func moveSecodeLargestValueToEnd(slice []int) {
	var num = len(slice)
	var sliceCopy = make([]int, num)
	var bubbleSortdLargestData = make([]int, num)
	copy(sliceCopy, slice)
	copy(bubbleSortdLargestData, bubbleSort(slice))
	for k, v := range slice {
		if bubbleSortdLargestData[num-2] == v {
			sliceCopy[k], sliceCopy[num-2] = sliceCopy[num-2], sliceCopy[k]
		}
	}
	fmt.Println("倒数第二大的值与末尾值交换位置后切片列表：", sliceCopy)
}

// 冒泡排序
func bubbleSort(slice []int) (data []int) {
	var num = len(slice)
	data = make([]int, num)
	copy(data, slice)
	for i := 0; i < num; i++ {
		for k, _ := range data[0:(num - i - 1)] {
			if data[k] > data[k+1] {
				data[k], data[k+1] = data[k+1], data[k]
			}
		}
	}
	return data
}

func main() {
	var slice = []int{9, 8, 19, 10, 2, 8}
	fmt.Printf("切片值：%v\n\n", slice)

	// 找出最大值以及索引
	findLargestValue(slice)
	fmt.Println("")

	// 找出第二大的值以及索引
	findSecodeValue(slice)
	fmt.Println("")

	// 最大值与末尾值交换位置
	moveLargestValueToEnd(slice)
	fmt.Println("")

	// 倒数第二大的值与末尾值交换位置
	moveSecodeLargestValueToEnd(slice)
	fmt.Println("")

	// 冒泡排序
	fmt.Println("冒泡排序后顺序：", bubbleSort(slice))
}
