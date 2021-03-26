package main

import (
	"fmt"
)

func main() {
	slice := []int{9, 8, 9, 19, 10, 2, 8}
	var s int
	var flag int
	fmt.Printf("请输入你要查找的数字：")
	fmt.Scan(&s)
	for i := 0; i < len(slice); i++ {
		switch s {
		case slice[i]:
			fmt.Printf("数字对应的索引为：%d\n", i)
			flag = 1
		default:
			flag = -1
		}
	}

	fmt.Println(flag)
	if flag == -1 {
		fmt.Printf("查找的数字不存在，返回值为：%d\n", -1)
	}
}
