package main

import "fmt"

func main() {
	nums := []int{6, 4, 5, 3, 1}
	// 将第一个最大的移动到最后
	// 0: 0(6) 1(4) true
	// 0: 4 6 5 3 1
	// 1: 1(6) 2(5) true
	// 1: 4 5 6 3 1
	// 2: 2(6) 3(3) true
	// 2：4 5 3 6 1
	// 3: 3(6) 4(1) true
	// 3: 4 5 3 1 6
	length := len(nums)
	for j := 0; j < length-1; j++ {
		for i := 0; i < length-1-j; i++ {
			// fmt.Printf("比较%d(%d)和%d(%d), 交换:%t\n", i, nums[i], i+1, nums[i+1], nums[i] > nums[i+1])
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
			// fmt.Printf("第%d次执行完成结果: %v\n", i, nums)
		}
		fmt.Println(nums)
	}
}
