package main

import "fmt"

func main() {
	slice_data := []int{9, 8, 19, 10, 2, 8}
	//for loop the total length
	for i := 1; i < len(slice_data); i++ {
		//Compare the adjacent elements
		for j := 0; j < (len(slice_data) - i); j++ {
			//IF the left of element > the right of element, and exchange the position
			if slice_data[j] > slice_data[j+1] {
				slice_data[j], slice_data[j+1] = slice_data[j+1], slice_data[j]
			}
		}
	}
	fmt.Println(slice_data)
}
