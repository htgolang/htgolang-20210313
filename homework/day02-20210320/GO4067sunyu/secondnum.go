package main

import "fmt"

func main() {
	slice := []int{9, 8, 19, 10, 2, 8}
	num := 0
	index := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > num {
			num = slice[i]
			index = i
		}
	}
	index1 := index + 1
	slice = append(slice[0:index], slice[index1:]...)
	num = 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > num {
			num = slice[i]
		}
	}
	fmt.Println("第二大数字是:", slice[index])
}
