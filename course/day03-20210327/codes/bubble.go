package main

import "fmt"

func bubble(nums [][2]int) {
	length := len(nums)
	for j := 0; j < length-1; j++ {
		for i := 0; i < length-1-j; i++ {
			// fmt.Printf("比较%d(%d)和%d(%d), 交换:%t\n", i, nums[i], i+1, nums[i+1], nums[i] > nums[i+1])
			if nums[i][1] > nums[i+1][1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
			}
			// fmt.Printf("第%d次执行完成结果: %v\n", i, nums)
		}
		// fmt.Println(nums)
	}
}

func main() {
	nums := [][2]int{{1, 6}, {2, 4}, {3, 9}, {4, 8}, {5, 1}}
	// nums排序
	// 按照每个元素索引为1的值比较 从小到大排序

	bubble(nums)
	fmt.Println(nums)
}
