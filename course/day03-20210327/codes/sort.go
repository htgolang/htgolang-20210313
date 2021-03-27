package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{6, 7, 9, 3, 10, 3, 8, 1}
	sort.Ints(nums)
	fmt.Println(nums)

	scores := []float64{1.2, 1.9, 2.0, 1.8, 1.1}
	sort.Float64s(scores)
	fmt.Println(scores)

	slice := [][2]int{{1, 2}, {2, 3}, {3, 6}, {4, 1}, {5. - 10}}
	less := func(i, j int) bool {
		// i < j
		return slice[i][1] < slice[j][1]
	}
	sort.Slice(slice, less)
	fmt.Println(slice)
}
