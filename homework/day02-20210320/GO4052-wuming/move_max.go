
package main

import "fmt"

func main() {

	forSlice := []int{9,8,19,10,2,8}


	maxNum := 0
	numIndex := 0

	for index, num := range forSlice{
		if num > maxNum {
			maxNum = num
			numIndex = index
		} else{
			continue
		}
	}
	fmt.Printf("原切片是 %v\n",forSlice)
	fmt.Printf("最大的数字是 %d,索引是 %d\n",maxNum,numIndex)
	tmp := append(forSlice[:numIndex],forSlice[numIndex+1:]...)
	slice := append(tmp,maxNum)
	fmt.Printf("调整后切片是 %v\n",slice)
}
