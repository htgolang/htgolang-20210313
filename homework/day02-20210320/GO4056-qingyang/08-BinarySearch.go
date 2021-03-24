package main

import "fmt"

func BinarySearch(slice []int, target, low, high int) int {
	if high < low {
		return -1
	}

	mid := (low + high) / 2
	if slice[mid] < target {
		return BinarySearch(slice, target, mid+1, high)
	} else if slice[mid] > target {
		return BinarySearch(slice, target, low, mid-1)
	}
	return mid
}

func main() {
	var target int
	slice := []int{2, 8, 9, 10, 19}
	fmt.Print("plz input a number: ")
	fmt.Scanln(&target)
	fmt.Println(BinarySearch(slice, target, 0, len(slice)-1))
}
