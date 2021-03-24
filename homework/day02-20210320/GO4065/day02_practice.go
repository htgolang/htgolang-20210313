package main

import (
	"fmt"
	"io/ioutil"
	"sort"
)

// CountText is a function to count the number of letters in text
func CountText() {
	var s = make(map[string]int)
	ctx, err := ioutil.ReadFile(`I_hava_a_dream.txt`)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(ctx); i++ {
		if (byte(ctx[i]) >= 65 && byte(ctx[i]) <= 90) || (byte(ctx[i]) >= 97 && byte(ctx[i]) <= 122) {
			var index string = string(byte(ctx[i]))
			v, ok := s[index]
			if ok {
				s[index] = v + 1
			} else {
				s[index] = 1
			}
		}
	}
	j := 0
	keys := make([]string, len(s))
	for k := range s {
		keys[j] = k
		j++
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, s[k])
	}
}

// FindMax is a function to find the largest value and return index.
func FindMax(arr []int) int {
	maxIndex, maxValue := 0, arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
			maxIndex = i
		}
	}
	return maxIndex
}

// FindSecond is a function to find the second largest value and return secondIndex.
func FindSecond(arr []int) int {
	maxValue := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
	}
	secondIndex, secondValue := 0, arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > secondValue && arr[i] != maxValue {
			secondValue = arr[i]
			secondIndex = i
		}
	}
	return secondIndex
}

// Max2End is a function
func Max2End(arr []int) []int {
	maxIndex := FindMax(arr)                                        // find the largest value's index
	arr[maxIndex], arr[len(arr)-1] = arr[len(arr)-1], arr[maxIndex] // exchange the two value
	return arr
}

// Second2Second is a function
func Second2Second(arr []int) []int {
	secondIndex := FindSecond(arr)                                        // find the second largest value's index
	arr[secondIndex], arr[len(arr)-2] = arr[len(arr)-2], arr[secondIndex] // exchange the two value
	return arr
}

// BubbleSort is a function to bubble sort the data.
func BubbleSort(arr []int) []int {
	fmt.Println("Before sorting is: ", arr)
	// Sort from small to big
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j <= i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr

	// Sort from big to small
	/*
		for i := 0; i <= len(arr)-1; i++ {
			for j := 0; j <= i-1; j++ {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
		return arr
	*/
}

// BinSearch is a function
func BinSearch(arr []int) int {
	var inputValue int
	fmt.Printf("please input any number in the arrary %d : ", arr)
	fmt.Scan(&inputValue)
	startIndex := 0
	endIndex := len(arr) - 1
	for startIndex <= endIndex {
		median := (startIndex + endIndex) / 2
		if arr[median] < inputValue {
			startIndex = median + 1
		} else {
			endIndex = median - 1
		}
	}
	if startIndex == len(arr) || arr[startIndex] != inputValue {
		return -1
	} else {
		// fmt.Printf("The index of the value %d is %d .\n", inputValue, startIndex)
		return startIndex
	}

}

func main() {
	// 统计打印出文本中大小字母出现的次数
	CountText()

	data := []int{9, 8, 19, 10, 2, 8}

	// 打印数组中最大值及其索引
	maxIndex := FindMax(data)
	fmt.Printf("The index of the largest value in the slice is %d, and its value is %d .\n", maxIndex, data[maxIndex])

	// 打印数组中第二大值及其索引
	secondIndex := FindSecond(data)
	fmt.Printf("The index of the second largest value in the slice is %d, and its value is %d .\n", secondIndex, data[secondIndex])

	// 把数组中最大值移至数组最后
	fmt.Printf("The result is: %d \n", Max2End(data))

	// 把数组中第二大值移至倒数第二位
	fmt.Printf("The result is: %d \n", Second2Second(data))

	// 冒泡排序，从小到大排序 or 从大到小排序
	fmt.Println("After sorting is: ", BubbleSort(data))

	// 二分查找法，打印出元素对应的索引
	data1 := []int{2, 8, 9, 10, 19}
	index := BinSearch(data1)
	fmt.Printf("The index of the value %d is %d . \n", data[index], index) // 输入数组内的值
	fmt.Println(BinSearch(data))                                           // 输入非数组内的值

}
