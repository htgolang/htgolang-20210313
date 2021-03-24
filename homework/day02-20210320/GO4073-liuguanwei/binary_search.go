package main

import (
	"fmt"
	"os"
)

func main() {
	slice_data := []int{2, 8, 9, 10, 19}
	left := 0
	right := len(slice_data)
	var num int
	fmt.Println("Pls enter a number[2, 8, 9, 10, 19]: ")
	fmt.Scan(&num)
	for left < right {
		//Define the middle position
		middle_index := left + (right-left)/2
		//If hit it fist time, then print result
		if num == slice_data[middle_index] {
			fmt.Println("The index is: ", middle_index)
			os.Exit(1)
		}
		//If the number smaller than the middle, move right side to middle side
		if num < slice_data[middle_index] {
			right = middle_index
			//If the number bigger than the middle, move forward the left side
		} else {
			left = middle_index + 1
		}
	}
	//If search failed, then print -1
	fmt.Println(-1)
}
