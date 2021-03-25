//Author:LJF
//Date:2021-03-24
package main

import "fmt"

func main() {
	//定义切片
	num := []int{9, 8, 19, 10, 2, 8}
	//假设切片的第一个元素为最大值
	maxNum := num[0]
	maxIndex := 0
	// secNum := num[0]
	// secIndex := 0
	//遍历切片中所有的元素，与第一个元素进行比较，如果大于第一元素就重新赋值并取其对应的索引
	for i := 1; i < len(num); i++ {
		if maxNum < num[i] {
			maxNum = num[i]
			maxIndex = i
		}

	}

	//通过切片方式获取然后组成新的数组
	fmt.Println("原切片为：", num)
	var num1 = num[:maxIndex]
	var num2 = num[maxIndex+1:]
	var num3 = num[maxIndex]
	//通过append组成新的切片
	var s1 = append(append(num1, num2...), num3)

	// fmt.Println("原切片为：", num)
	fmt.Println("新的切片为：", s1)

}

// 打印结果为：
// 原切片为： [9 8 19 10 2 8]
// 新的切片为： [9 8 10 2 8 19]
