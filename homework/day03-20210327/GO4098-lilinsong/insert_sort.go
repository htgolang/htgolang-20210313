package main

import (
	"fmt"
	"math/rand"
	"time"
)

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1
		for j >= 0 && tmp < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = tmp
	}
}

func main() {
	var arr = make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		randNum := rand.Intn(200)
		arr[i] = randNum
	}
	fmt.Printf("1, %p, %v\n", arr, arr)
	insertSort(arr)
	fmt.Printf("3, %p, %v\n", arr, arr)
}
