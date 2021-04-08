package main

import "fmt"

func InsertSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		tmp := arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] > tmp {
				arr[j+1], arr[j] = arr[j], tmp
			}
		}
	}
}

func main() {
	var arr = []int{0, 5, 8, 3, 2, 4, 1}
	InsertSort(arr)
	fmt.Println(arr)
}
