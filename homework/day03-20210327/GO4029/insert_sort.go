package main

import "fmt"

func main() {
	var arr = [7]int{2, 10, 3, 5, 6, 4, 7}
	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}
	fmt.Printf("%v", arr)
}
