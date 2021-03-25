package main

import "fmt"

func main() {

	gSlice := []int{2, 5, 9, 10, 19}
	var input int
	fmt.Printf("请输入要查找的数字：")
	fmt.Scan(&input)
	low := 0
	high := len(gSlice) - 1
	for low <= high {
		mid := (low + high) / 2
		if gSlice[mid] == input {
			fmt.Printf("%v", mid)
			return
		} else if gSlice[mid] < input {
			low = mid + 1
		} else if gSlice[mid] > input {
			high = mid - 1
		}
	}
	fmt.Printf("-1")
}
