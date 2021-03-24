package main

import "fmt"

// 参考：https://leetcode-cn.com/problems/pascals-triangle-ii/solution/zhi-xing-yong-shi-0-ms-zai-suo-you-golang-ti-jia-8/

func getRow(num int) []int {
	nums := []int{1}
	for i := 1; i <= num; i++ {
		// 尾部追加1
		nums = append(nums, 1)
		// 倒序计算杨辉三角当前行
		for j := i - 1; j > 0; j-- {
			nums[j] += nums[j-1]
		}
	}
	return nums
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(getRow(i))
	}
}
