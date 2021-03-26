package main

import "fmt"

func searchSecondNum(arr []int) {
	var (
		max, min, index int
	)
	for i, v := range arr {
		if v > max {
			max = v
		} else if min < max && min < v {
			min = v
			index = i
		}
	}
	fmt.Println(index, min)
}

func main() {
	arr := []int{9, 8, 19, 10, 2, 8}
	searchSecondNum(arr)
}
