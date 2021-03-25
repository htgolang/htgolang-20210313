//Author:LJF
//Date:2021-03-24
package main

import "fmt"

func main() {
	//定义切片
	num := []int{9, 8, 19, 10, 2, 8}
	//假设第一个元素为最大值
	maxNum := num[0]
	maxIndex := 0
	secNum := num[0]
	secIndex := 0
	//遍历切片中所有的元素，与第一个元素进行比较，如果大于第一元素就重新赋值并取其对应的索引
	for i := 1; i < len(num); i++ {
		if maxNum < num[i] {
			maxNum = num[i]
			maxIndex = i
		}

	}
	for i := 1; i < len(num); i++ {
		if num[i] > secNum {
			if num[i] < maxNum {
				secNum = num[i]
				secIndex = i
			}

		}

	}
	fmt.Printf("最大的元素的索引为：%v，最大的元素的值为：%v\n", maxIndex, maxNum)
	fmt.Printf("第二大的元素的索引为：%v，第二大的元素的值为：%v\n", secIndex, secNum)
}
