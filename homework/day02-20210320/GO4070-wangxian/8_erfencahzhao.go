package main

import "fmt"

func main() {
	var s1 = []int{1, 3, 8, 9, 11, 14, 22, 34, 35}

	var target_num int

	fmt.Print("请输入你要查找的数字：")
	fmt.Scanln(&target_num)

	var left = 0
	var right = len(s1) -1
	var flag bool

	for left <= right {
		mid := (left + right) /2

		if s1[mid] == target_num{
			fmt.Printf("存在，索引是%d\n", mid)
			flag = true
			break
		} else if s1[mid] > target_num{
			right = mid - 1
		} else if s1[mid] < target_num{
			left = mid + 1
		}
	}

	if !flag{
		fmt.Println("-1")
	}
}
