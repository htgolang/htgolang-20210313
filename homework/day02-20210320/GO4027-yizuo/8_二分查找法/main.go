package main

import "fmt"

// 用户字符串输入返回
func userStrInputReturn(str string) (inputValue string) {
	fmt.Println(str)
	fmt.Scan(&inputValue)
	return inputValue
}

// 用户数字输入返回
func userIntInputReturn(str string) (inputValue int) {
	fmt.Println(str)
	fmt.Scanln(&inputValue)
	return inputValue
}

// 二分法查找
func binarySearchMethod(slice []int, inputValue int) {
	/*
		算法描述：在一组有序数组中，将数组一分为二，将要查询的元素和分割点进行比较，分为三种情况
			相等直接返回
			元素大于分割点，在分割点右侧继续查找
			元素小于分割点，在分割点左侧继续查找
	*/
	left, right, mid := 0, len(slice), 0

	for {
		mid = (left + right) / 2

		// 如果左侧大于右侧则，应该报错退出
		// 如果中间值大于整体长度，应该报错退出
		if left > right || mid >= len(slice) {
			fmt.Println("输入的值在列表中不存在！")
			break
		}

		// 如果中间值大于输入值，则右侧应该为中间值减1
		// 如果中间值小于输入值，则左侧应该为中间值加1
		// 如刚好中间值等于输入值，那么列表中存在
		if slice[mid] > inputValue {
			right = mid - 1
		} else if slice[mid] < inputValue {
			left = mid + 1
		} else {
			fmt.Println("值存在，位置为: ", mid)
			break
		}
	}

}

func main() {
	for {
		// 输入查找逻辑
		slice := []int{2, 8, 9, 10, 19}
		userIntValue := userIntInputReturn("请输入你要查询的值: ")
		binarySearchMethod(slice, userIntValue)

		// 确认是否继续逻辑
		userStrValue := userStrInputReturn("是否继续查找，继续请输入yes，默认退出: ")
		if userStrValue != "yes" {
			break
		}
	}

}
