package main

import "fmt"

func maxNumEnd(arr []int, left, right int) []int {
	if left < right {
		mid := partition(arr, left, right)
		maxNumEnd(arr, left, mid-1)
		maxNumEnd(arr, mid+1, right)
	}
	return arr
}

func partition(arr []int, left, right int) int {
	tmp := arr[left]
	for left < right {
		for left < right && arr[right] >= tmp {
			right--
		}
		arr[left] = arr[right]

		for left < right && arr[left] <= tmp {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = tmp

	return left
}

func main() {
	arr := []int{9, 8, 19, 10, 2, 8}
	ret := maxNumEnd(arr, 0, len(arr)-1)
	fmt.Printf("第二大的值是: %d\n", ret[len(arr)-2])

}
