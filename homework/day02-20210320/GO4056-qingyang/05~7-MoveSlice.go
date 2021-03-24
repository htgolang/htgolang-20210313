package main

import "fmt"

// 5,6,7用冒泡可以全部完成
func BubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 1; j < len(slice)-i; j++ {
			if slice[j] < slice[j-1] {
				slice[j], slice[j-1] = slice[j-1], slice[j]
			}
		}
	}
}

func main() {
	slice := []int{9, 8, 19, 10, 2, 8}
	BubbleSort(slice)
	fmt.Println(slice)
}
