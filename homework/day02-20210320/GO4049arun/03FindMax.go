package main

import "fmt"

func main() {
	var data = []int{9, 8, 19, 10, 2, 8}
	var maxIndex int = 0
	var maxValue int = data[0]
	for i := 0; i < len(data); i++ {
		if data[i] > maxValue {
			maxValue = data[i]
			maxIndex = i
		}
	}
	fmt.Printf("切片中最大的值为:%d,最大的索引为:%d\n", maxValue, maxIndex)
}

/*
> go run .\03FindMax.go
切片中最大的值为:19,最大的索引为:2
*/
