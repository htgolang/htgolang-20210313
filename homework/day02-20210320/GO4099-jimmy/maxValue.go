package main

import "fmt"

func main() {
	var max, index int

	nums := []int{9, 8, 19, 10, 2, 8}
	for i, v := range nums {
		if v > max {
			max = v
			index = i
		}
	}
	fmt.Println(index, max)
}
