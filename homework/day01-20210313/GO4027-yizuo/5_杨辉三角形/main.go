package main

import (
	"fmt"
)

/*
　　　　　　　　１
　　　　　　　１　１
　　　　　　１　２　１
　　　　　１　３　３　１
　　　　１　４　６　４　１
　　　１　５　10　10　５　１
　　１　６　15　20　15　６　１
　１　７　21　35　35　21　７　１
１　８　28　56　70　56　28　８　１

百科：
杨辉三角形，又称帕斯卡三角形、贾宪三角形、海亚姆三角形、巴斯卡三角形，是二项式系数的一种写法，形似三角形，在中国首现于南宋杨辉的《详解九章算法》得名，书中杨辉说明是引自贾宪的《释锁算书》，故又名贾宪三角形。前 9 行写出来如上：

解析：
1、第一个值与最后一个值一定是1
2、有多少行，那一行就有多少个值
3、除了最后一个值，其他的值都是上级同位置的值加位置-1的值。
*/

// 杨辉三角

// 输入打印行数
func accessInputNumber() (inputNumber int) {
	fmt.Println("请输入你想打印杨惠三角形的行数:")
	fmt.Scan(&inputNumber)

	return inputNumber
}

//
func binomialArray() {
	inputNumber := accessInputNumber()
	binArray := []int{}

	for i := 0; i < inputNumber; i++ {

		for j := 0; j < (i + 1); j++ {
			var length = len(binArray)
			var value int

			if j == 0 || j == i {
				value = 1
			} else {
				value = binArray[length-i] + binArray[length-i-1]
			}

			binArray = append(binArray, value)
			fmt.Print(value, " ")
		}
		fmt.Println("")
	}
}

func main() {
	binomialArray()
}
