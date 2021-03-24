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
	tmp := data[len(data)-1]
	data[len(data)-1] = maxValue
	data[maxIndex] = tmp
	fmt.Println(data) //[9 8 8 10 2 19]

	var secondIndex int = 0
	var secondValue int = data[0]
	for i := 0; i < len(data); i++ {
		if data[i] > secondValue && data[i] != maxValue {
			secondValue = data[i]
			secondIndex = i
		}
	}
	tmp = data[len(data)-2]
	data[len(data)-2] = secondValue
	data[secondIndex] = tmp
	fmt.Println(data)
}
