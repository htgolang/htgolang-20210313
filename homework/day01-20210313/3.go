package main

import "fmt"

func main() {
	var n int       //储存打印的行数
	nums := []int{} //储存每一行的数字
	fmt.Println("请输入要打印的行数:")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < (n - i); j++ {
			fmt.Print(" ")
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
			fmt.Printf("%v ", value)
		}
		fmt.Println("")
	}

}
