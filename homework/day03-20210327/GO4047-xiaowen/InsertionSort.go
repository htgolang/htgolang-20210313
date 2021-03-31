package main

import "fmt"

func main() {
	arr := []int{3,76,23,3,65,13,65,8,1}
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
				arr[preIndex+1] = arr[preIndex]
				preIndex -= 1
		}
		arr[preIndex+1] = current
	}
	fmt.Println(arr) 
}