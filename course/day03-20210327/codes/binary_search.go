package main

import "fmt"

func main() {
	nums := []int{2, 8, 9, 10, 19}

	num := 20
	middle := 0

	start, end := 0, len(nums)-1
	for {
		middle = (start + end) / 2 //2

		if nums[middle] == num {
			fmt.Println("找到了", middle)
			break
		} else if nums[middle] > num { // 在数组前半段找
			end = middle - 1
		} else {
			start = middle + 1
		}
		if start > end {
			fmt.Println("未找到")
			break
		}
	}
}
