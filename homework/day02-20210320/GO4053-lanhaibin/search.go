package main

import "fmt"

func main() {
	slice := []int{9, 8, 19, 10, 2, 8}
	fmt.Println("初始切片:")
	fmt.Println(slice)
	var maxIndex int = 0
	for i:=1;i<len(slice);i++ {
		if slice[i] > slice[maxIndex] {
			maxIndex = i
		}
	}
	fmt.Printf("最大的数是: %d\n", slice[maxIndex])
	fmt.Println("移动最大的数到末尾.")
	slice[maxIndex], slice[len(slice)-1] = slice[len(slice)-1], slice[maxIndex]
	fmt.Println(slice)

	var secondIndex int = 0
	for i:=1;i<len(slice)-1;i++ {
		if slice[i] > slice[secondIndex] {
			secondIndex = i
		}
	} 
	fmt.Printf("第二大的数是: %d\n", slice[secondIndex])
	fmt.Println("移动第二大的数到倒数第二")
	slice[secondIndex], slice[len(slice)-2] = slice[len(slice)-2], slice[secondIndex]
	fmt.Println(slice)
}