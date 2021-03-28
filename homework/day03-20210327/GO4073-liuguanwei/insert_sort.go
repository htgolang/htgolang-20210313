package main

import "fmt"

func main() {
	slice_data := []int{9, 8, 19, 10, 2, 8, 1}
	//Start from the second number
	for i := 1; i < len(slice_data); i++ {
		//Get the value of every index
		value := slice_data[i]
		//Index from the left side, compare the number from left side
		index := i - 1
		//It will at the beginning of the map, when index == 0
		for index >= 0 {
			//index is getting smaller
			if value < slice_data[index] {
				slice_data[index+1] = slice_data[index] //Shift number from slot i to slot i+1
				slice_data[index] = value               //shift value left int slot
				index = index - 1
			} else {
				break
			}
		}

	}
	fmt.Println(slice_data)
}
