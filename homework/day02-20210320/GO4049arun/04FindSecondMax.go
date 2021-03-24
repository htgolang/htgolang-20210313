package main

import "fmt"

func main() {
	var data = []int{9, 8, 19, 10, 2, 8}
	var maxValue int = data[0]
	for i := 0; i < len(data); i++ {
		if data[i] > maxValue {
			maxValue = data[i]
		}
	}
	var secondIndex int = 0
	var secondValue int = data[0]
	for i := 0; i < len(data); i++ {
		if data[i] > secondValue && data[i] != maxValue {
			secondValue = data[i]
			secondIndex = i
		}
	}

	fmt.Printf("切片中第2大的值为:%d,第2大的索引为:%d\n", secondValue, secondIndex)
}

/*
> go run .\04FindSecondMax.go
切片中第2大的值为:10,第2大的索引为:3
*/
