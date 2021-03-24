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

	fmt.Println(data)
}

/*
> go run .\05migrateMax2End.go
[9 8 8 10 2 19]
*/
