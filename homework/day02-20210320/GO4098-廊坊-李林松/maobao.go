package main

import "fmt"

func maoBao(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i ; j < len(arr)-1; j++ {
			if arr[i] > arr[j+1] {
				arr[i], arr[j+1] = arr[j+1], arr[i]
			}
		}
	}
	fmt.Println(arr)
}
func main() {
	arr := []int{9, 8, 19, 10, 2, 8}
	maoBao(arr)
}
