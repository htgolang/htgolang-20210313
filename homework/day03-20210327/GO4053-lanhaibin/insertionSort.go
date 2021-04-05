package main

import "fmt"

func insertionSort(slice []int) {
	for i:=1;i<len(slice);i++ {
		v := slice[i]
		var j int
		for j=i;j>0;j-- {
			if v < slice[j-1] {
				slice[j] = slice[j-1]
			} else {
				break
			}
		}
		slice[j] = v
	}
}

func main() {
	slice := []int{8,3,6,1,8,3,6,8,3,2,6,4,9}
	insertionSort(slice)
	fmt.Println(slice)
}