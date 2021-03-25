package main

import (
	"fmt"
)

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
	slice1 := append(slice[0:index], slice[index1:]...)
	num1 := 0
	for i := 0; i < len(slice1); i++ {
		if slice1[i] > num1 {
			num1 = slice1[i]
			index1 = i
		}
	}
	fmt.Println(slice, num1, index, slice1)
	init := []int{9, 8, 19, 10, 2, 8}
	tmp := init[len(init)-2]
	fmt.Println(tmp)
	for k, _ := range init {
		if init[k] == num1 {
			init[k] = tmp
		}
	}
	init[len(init)-2] = num1
	fmt.Println(init)

	//	init[len(init)-2] = num1

	//	fmt.Println(init)

}
