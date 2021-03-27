package main

import "fmt"

func main() {
	// 匿名函数
	// 匿名结构体
	// 匿名接口
	// 没有名字的函数
	// 定义后直接赋值给某个变量 直接调用
	// 辅助性的功能
	add := func(left, right int) int {
		return left + right
	}

	fmt.Printf("%T %v\n", add, add)

	fmt.Println(add(1, 2))

	nums := []int{1, 2, 3, 4, 5}
	// mapf => 新[]int, tl(int) int

	tlAdd5 := func(n int) int {
		return n + 5
	}
	fmt.Println(mapf(nums, tlAdd5))

	x := func(left, right int) int {
		return left + right
	}(1, 2)

	fmt.Println(x)

	func(left, right int) int {
		fmt.Println("lambda call")
		return left + right
	}(1, 2)

	fmt.Println(func(left, right int) int {
		return left + right
	}(1, 2))

}

func mapf(list []int, tl func(int) int) []int {
	rt := []int{}
	for _, e := range list {
		rt = append(rt, tl(e))
	}
	return rt
}
