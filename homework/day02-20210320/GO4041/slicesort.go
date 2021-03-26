package main

import "fmt"

func main() {
	nums := []int{9, 8, 19, 10, 2, 8}
	fmt.Println("原始切片：", nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			nums[i+1], nums[i] = nums[i], nums[i+1]
		}
	}
	fmt.Println("最大值移动到末尾", nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > nums[i+1] {
			nums[i+1], nums[i] = nums[i], nums[i+1]
		}
	}
	fmt.Println("第二大值移动到倒数第二", nums)

	fmt.Println("为了显示效果切片重置")
	nums = []int{9, 8, 19, 10, 2, 8}
	fmt.Println("原始切片：", nums)
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
		fmt.Println("冒泡过程", nums)
	}
	fmt.Println("冒泡排序", nums)

	var in int
	fmt.Println("请输入您需要查找的值：")
	fmt.Scan(&in)

	f := false
	for j, i := range nums {
		if i == in {
			fmt.Printf("存在，索引：%d\n", j)
			f = true
			break
		}
	}
	if !f {
		fmt.Printf("不存在\n")
	}

}
