package main

import "fmt"

func bubble(nums []int) {
	length := len(nums)
	for j := 0; j < length-1; j++ {
		for i := 0; i < length-1-j; i++ {
			// fmt.Printf("比较%d(%d)和%d(%d), 交换:%t\n", i, nums[i], i+1, nums[i+1], nums[i] > nums[i+1])
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
			// fmt.Printf("第%d次执行完成结果: %v\n", i, nums)
		}
		// fmt.Println(nums)
	}
}

func main() {
	nums := []int{6, 4, 5, 3, 1}

	bubble(nums)
	fmt.Println(nums)

	nums = []int{6, 4000, 5000, 300, 100}

	bubble(nums)
	fmt.Println(nums)
}
