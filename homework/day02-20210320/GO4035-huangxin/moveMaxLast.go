package main

/*
	create-date:2021-03-21
	describe: 移动最大值到末尾
	author:GO4035-huangxin
*/

import "fmt"

func main() {
	/*
			将最大的元素移动到切片最末尾
		    []int{9, 8, 19, 10, 2, 8}
	*/
	nums := []int{9, 8, 19, 10, 2, 8}

	// 找到 最大值 与其 index
	maxIndex, maxNum := getMaxNum(nums)

	// copy() 删除最大值
	copy(nums[maxIndex:], nums[maxIndex+1:])

	// 更新nums
	nums = nums[:len(nums)-1]
	// 将最大值 追加到 nums
	nums = append(nums, maxNum)

	// 打印输出 nums
	fmt.Println(nums)

	// fmt.Println(maxNum, maxIndex)
}

// 返回最大值 与其 index
func getMaxNum(nums []int) (int, int) {

	max_num := 0
	max_index := 0

	for i, v := range nums {
		if max_num < v {
			max_num = v
			max_index = i
		}
	}

	return max_index, max_num
}
