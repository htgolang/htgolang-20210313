package main

import "fmt"

func FindMaxInSlice(slice []int) (int, int) {
	max := slice[0]
	maxIndex := 0
	for i, v := range slice {
		if v > max {
			max, maxIndex = v, i
		}
	}
	return maxIndex, max
}

func main() {
	index, max := FindMaxInSlice([]int{9, 8, 19, 10, 2, 8})
	fmt.Println(index, max)
}
