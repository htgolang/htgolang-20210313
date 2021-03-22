/**
* @Author: fengzhilaoling
* @Date: 2021/3/21 8:57
**/
package main

import "fmt"

func main() {
	var nums = []int{9, 8, 19, 10, 2, 8}
	// 冒泡排序
	nums1 := testTwo5(nums)
	// 最大值的索引
	maxIndex, _ := testTwo1(nums1[len(nums1)-1], nums)
	// 第二大的索引
	max2Index, _ := testTwo1(nums1[len(nums1)-1], nums)
	// 最大值移动到最后
	nums[maxIndex], nums[len(nums)-1] = nums[maxIndex], nums[len(nums)-1]
	fmt.Println(nums)
	// 第二大值移动到最后
	nums[max2Index], nums[len(nums)-2] = nums[max2Index], nums[len(nums)-2]
	fmt.Println(nums)

}

// 返回最大的值和的索引
func testTwo1(maxNum int, nums []int) (int, int) {
	for index, v := range nums {
		if v == maxNum {
			fmt.Printf("nums切片中最大元素的索引是：%d,最大值为：%d\n", index, v)
			return index, v
		}
	}
	return 0, 0
}

// 返回元素中第二大的值和索引
func testTwo2(maxNum2 int, nums []int) (int, int) {
	for index, v := range nums {
		if v == maxNum2 {
			fmt.Printf("nums切片中最大元素的索引是：%d,最大值为：%d\n", index, v)
			return index, v
		}
	}
	return 0, 0
}

// 冒泡排序
func testTwo5(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	fmt.Println("排序后的切片：")
	fmt.Println(nums)
	return nums
}
