package main

import "fmt"

func main() {

	forSlice := []int{9,8,19,10,2,8}


	maxNum := 0
	secondNum := 0
	numIndex := 0

	for index, num := range forSlice{
		if num > maxNum {
			maxNum,secondNum = num,maxNum
		} else if num < maxNum  {
			if num > secondNum{
				secondNum = num
				numIndex = index
			}
		}else{
			continue
		}
	}
	fmt.Printf("原切片是 %v\n",forSlice)
	fmt.Printf("第二大的数字是 %d,索引是 %d\n",secondNum,numIndex)
	tmp := append(forSlice[:numIndex],forSlice[numIndex+1:]...)
	tmp2 :=tmp[:len(tmp)-1]
	lastnum :=tmp[len(tmp)-1]
	slice :=append(tmp2,secondNum,lastnum)
	fmt.Printf("调整后切片是 %v\n",slice)
}
