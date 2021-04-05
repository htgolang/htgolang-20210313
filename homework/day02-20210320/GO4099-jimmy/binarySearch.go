package main

import "fmt"

func main() {
	s := []int{2, 8, 9, 10, 19}
	var leftIndex, res int
	rightIndex := len(s) - 1
	fmt.Print("请输入一个整数[2, 8, 9, 10, 19]: ")
	fmt.Scan(&res)

	for leftIndex <= rightIndex {
		middle := (leftIndex + rightIndex) / 2
		if s[middle] > res {
			rightIndex = middle - 1
		} else if s[middle] < res {
			leftIndex = middle + 1
		} else {
			fmt.Println(middle)
			return
		}
	}
	fmt.Println("-1")
}
