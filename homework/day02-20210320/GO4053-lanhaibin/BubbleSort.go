package main

import "fmt"

func main() {
	slice := []int{9,8,19,10,2,8} 
	// fmt.Println("切片初始状态：")
	fmt.Println(slice)
	for i:=1;i<len(slice)-1;i++ {
		for j:=0;j<len(slice)-i;j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
	fmt.Println("排序后:")
	fmt.Println(slice)
}