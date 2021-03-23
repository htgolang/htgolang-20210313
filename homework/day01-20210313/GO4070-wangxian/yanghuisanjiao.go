package main

import "fmt"

func main() {
	//指定行数
	const lines int = 10

	nums := []int{}

	for i := 0; i < lines; i++ {
		for j := 0; j < (lines - i); j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}
			nums = append(nums, value)
			fmt.Print(value, " ")
		}
		fmt.Println()
	}
}
