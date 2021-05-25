package main

import "fmt"

func main() {
	i := 1
	a := [2]int{0, 1}
	s := []int{0, 1}

	fmt.Printf("%p\n", &i)
	fmt.Printf("%p %p\n", &a, &a[0])
	fmt.Printf("%p %p %p\n", &s, s, &s[0])

	// i
	// 数组 => 值类型
	// 切片 => 引用类型的 // sliceHeader{pointer, len, cap}

	paramsters[0]
}
