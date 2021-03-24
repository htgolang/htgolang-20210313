package main

import "fmt"

func FindSecondMaxInSlice(slice []int) (int, int) {
	max := slice[0]
	secondMax := 0
	maxIndex := 0
	secondIndex := 0
	for i, v := range slice {
		if v > max {
			secondMax = max
			secondIndex = maxIndex
			max = v
			maxIndex = i
		} else {
			if v > secondMax {
				secondMax = v
				secondIndex = i
			}
		}
	}
	return secondIndex, secondMax
}

func main() {
	index, max := FindSecondMaxInSlice([]int{9, 8, 19, 10, 2, 8})
	fmt.Println(index, max)
}
