package main

/*
	create-date:2021-03-21
	describe: 通过控制台输入，查找元素是否存在
	author:GO4035-huangxin
*/

import "fmt"

func main() {

	/*
			[]int{2, 8, 9, 10, 19}

		    控制台输入 9 => 查找在切片中的位置
		        不存在：-1
		        存在：索引
	*/

	// 初始化 nums
	nums := []int{2, 8, 9, 10, 19}

	fmt.Print("请输入一个数:")
	// 定义控制台输入
	var num int
	fmt.Scan(&num)

	// 定义接收 查询元素结果
	index := search(nums, num)
	if index != -1 {
		fmt.Printf("%d 该元素在nums中的索引为%d\n", num, index)
	} else {
		fmt.Printf("nums中不存在该元素%d\n", num)
	}

}

// 定义查找元素的方法
func search(nums []int, element int) int {

	var index int = -1
	for i, v := range nums {
		if v == element {
			index = i
			break
		}
	}
	return index
}
