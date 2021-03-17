package main

import "fmt"

var LINE int = 10

func main() {
	nums := []int{}
	for i := 0; i < LINE; i++ {
		for j := 0; j < (LINE - i); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < (i + 1); j++ {
			lenth := len(nums)
			var v int

			if j == 0 || j == i {
				v = 1
			} else {
				v = nums[lenth-i] + nums[lenth-i+1]
			}
			nums = append(nums, v)
			fmt.Print(v, " ")
		}
		fmt.Println("")
	}
}
