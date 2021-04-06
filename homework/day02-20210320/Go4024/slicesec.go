package main

import (
	"fmt"
)

func main() {
	var slice []int = []int{9, 8, 19, 10, 2, 8}
	var maxNum int
	var maxIndex int
	var secNum int
	var secIndex int

	for k,v := range slice {
		if	v > maxNum {
			maxNum = v
			maxIndex = k
		} else if v > secNum {
			secNum = v
			secIndex = k
		} else {
			continue
		}
	}
	fmt.Println("the Sec number is:",secNum,"the index is:",secIndex)
	fmt.Println("the Max number is:",maxNum,"the index is:",maxIndex)

	//移动最大的值到最后一位
	tmp := append(slice[:maxIndex],slice[maxIndex+1:]...)
	
	//fmt.Println(tmp)
	sliceMax := append(tmp,maxNum)
	//fmt.Println(slice)
	fmt.Println("移动最大到最后:",sliceMax)


	//移动第二大数字到倒数第二位
	slice = []int{9, 8, 19, 10, 2, 8}
	tmp1 := append(slice[:secIndex],slice[secIndex+1:]...)
	tmp2 := tmp1[:len(tmp1)-1]
	tmp3 := tmp1[len(tmp1)]
	slice = append(tmp2,secNum,tmp3)
	fmt.Println(slice)
}