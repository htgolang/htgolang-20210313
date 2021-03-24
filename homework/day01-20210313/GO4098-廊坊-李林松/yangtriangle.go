package main

import "fmt"

func yangTriangle(row int) [][]int {
	var triangle = make([][]int, row)
	for i := range triangle {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j] + triangle[i-1][j-1]
		}
	}
	return triangle

}

func main() {
	ret := yangTriangle(10)
	for _, v := range ret {
		fmt.Println(v)
	}
}
