package main

import "fmt"

func main() {
	// [length]T
	// T 可以是任意类型
	// T -> [5]int
	var multi [3][5]int // T = [5]int 3行5列的数组

	fmt.Printf("%T, %v\n", multi, multi)

	var multi3 [2][3][5]int
	fmt.Printf("%T, %v\n", multi3, multi3)

	multi2 := [...][5]int{
		{1, 3, 5, 6, 7},
		{},
		{1: 3, 3: 6, 4: 3},
	}

	fmt.Println(multi2)

	// 操作 == !=
	fmt.Println(multi == multi2)
	// 元素访问和修改
	fmt.Println(multi2)
	fmt.Println(multi2[1])
	fmt.Println(multi2[0][3])

	//
	multi2[1] = [5]int{1, 2, 3, 4, 5}
	fmt.Println(multi2[1])
	fmt.Println(multi2)

	multi2[0][3] = 1000
	fmt.Println(multi2)

	for i := 0; i < len(multi2); i++ {
		for j := 0; j < len(multi2[i]); j++ {
			fmt.Println(i, j, multi2[i][j])
		}
	}

	for index, array := range multi2 {
		fmt.Println(index, array)
		for j, value := range array {
			fmt.Println(index, j, value)
		}
	}
}
