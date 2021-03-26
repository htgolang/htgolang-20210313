package main

import "fmt"

func searchMaxNum(arr []int) {
	var max int
	var index int
	for i, v := range arr {
		if v > max {
			max = v
			index = i
		}
	}
	fmt.Printf("index=%d\tvalue=%d\n", index, max)
}

func main() {
	arr := []int{9, 8, 19, 10, 2, 8}
	searchMaxNum(arr)

}
