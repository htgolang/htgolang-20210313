package main

import "fmt"

const Liens int = 8

func ShowYhTriangle() {
	nums := []int{}
	for i := 0; i < Liens; i++ {
		//补空白
		for j := 0; j < (Liens - i); j++ {
			fmt.Println(" ")
		}
		for j := 0; j < (i + 1); j++ {
			//当前数组的长度
			var length = len(nums)
			//生成的数字
			var value int
			//每行开头和最后的数字为1
			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}
			nums = append(nums, value)
			fmt.Print(value, " ")
		}
		//换行
		fmt.Println("")
	}
}

func main() {
	ShowYhTriangle()
}
