package main

import "fmt"

func main() {
	slice := []int{2, 8, 9, 10, 19}
	fmt.Println(slice)
	var ans int
	fmt.Printf("请输入要查找的数: ")
	fmt.Scanf("%d", &ans)
	left, right := 0, len(slice)-1
	for left < right {
		mid := (left + right) / 2
		if slice[mid] == ans {
			fmt.Printf("查找成功，目标索引为: %d!\n", mid)
			break
		}
		if slice[mid] > ans {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left == right {
		fmt.Println("查找失败！")
	}
}