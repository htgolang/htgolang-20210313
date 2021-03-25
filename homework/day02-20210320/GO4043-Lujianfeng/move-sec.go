//Author:LJF
//Date:2021-03-24
/*
需求 找最大值
1. 循环,找到最大值
2. 进行比较最大值，找到第二大值
3.通过第二大值的索引截取新的切片
4. 通过append重新组合切片
5.通过索引交换第二大值
*/
package main

import "fmt"

func main() {
	//1.定义切片
	num := []int{9, 8, 19, 10, 2, 8}
	//假设数组的第一个元素为最大值
	maxNum := num[0]
	maxIndex := 0
	secNum := num[0]
	secIndex := 0
	//2.遍历切片中所有的元素，与第一个元素进行比较，如果大于第一元素就重新赋值并取其对应的索引
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

	//通过切片方式获取然后组成新的切片
	var num1 = num[:secIndex]
	var num2 = num[secIndex+1:]
	var num3 []int
	num3 = append(num3, num1...)
	num3 = append(num3, num2...)

	num3 = append(num3, secNum)
	num3[len(num3)-1], num3[len(num3)-2] = num3[len(num3)-2], num3[len(num3)-1]

	fmt.Printf("最大元素的索引为：%d, 最大元素的值为：%d\n", maxIndex, maxNum)
	fmt.Printf("第二大元素的索引为：%d, 第二大元素的值为：%d\n", secIndex, secNum)
	fmt.Println("原切片为：", num)
	fmt.Println("新的切片为：", num3)

}
