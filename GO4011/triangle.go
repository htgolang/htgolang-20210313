package main

import "fmt"

// 确认行数
const LINES int = 8

// 打印杨辉三角形
func main() {
	nums := []int{}

	for i := 0; i < LINES; i++ {
		for j := 0; j < (LINES - i); j++ {
			fmt.Print(" ")
		}

		for j := 0; j < (i + 1); j++ {
			// 当前数组长度
			var length = len(nums)
			// 本位置应该生成的数字
			var value int

			// 每行第一个元素和最后一个元素是1，中间元素=上方两个元素之和
			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}
			nums = append(nums, value)
			fmt.Print(value, " ")
		}
		fmt.Println("")
	}
}
