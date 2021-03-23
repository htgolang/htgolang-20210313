package main

import "fmt"

const Liens int = 8

func ShowYhTriangle() {
	nums := []int{}
	for i := 0; i < Liens; i++ {
		for j := 0; j < (Liens - i); j++ {
			fmt.Println(" ")
		}
		for j := 0; j < (i + 1); j++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}
			nums = append(nums, value)
			fmt.Print(value, "\t")
		}
		fmt.Println("")
	}
}

func main() {
	ShowYhTriangle()
}
