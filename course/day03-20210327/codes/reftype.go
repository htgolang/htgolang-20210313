package main

import "fmt"

func main() {
	nums := [5]int{1, 2, 3, 4, 5}
	nums2 := nums

	fmt.Println(nums, nums2)
	nums2[0] = 1000
	fmt.Println(nums, nums2)

}
