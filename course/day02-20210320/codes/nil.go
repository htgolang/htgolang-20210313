package main

import "fmt"

func main() {
	// 显示为切片和map进行初始化值
	var nilSlice []int
	var nilMap map[string]int
	// var nilSlice []int = make([]int, 0)
	// var nilMap map[string]int = make(map[string]int)

	fmt.Println(append(nilSlice, 1))
	fmt.Println(nilMap["kk"])
	nilMap["kk"] = 1
}
