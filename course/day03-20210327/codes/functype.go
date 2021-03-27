package main

import "fmt"

// 函数类型->函数的签名

func add(l, r int) int {
	return l + r
}

func sayHello() {

}

func main() {
	fmt.Printf("%T\n", add)

	// 定义函数类型的变量
	var f func(int, int) int // 零值 nil

	fmt.Printf("%T %v\n", f, f)

	f = add // 引用函数add

	fmt.Printf("%T %v\n", f, f)
	fmt.Println(f(2, 3)) // 调用函数

	// map, filter, reduce
	// 对切片进行操作
	// map: 对切片中每个元素通过某种转换得到结果组成的新切片
	// filter: 对切片中的元素进行过滤
	// reduce: 初始化元素 将1个元素与初始化元素 => result + 2 => result

	nums := []int{1, 2, 3, 4, 5}
	// mapf => 新[]int, tl(int) int

	fmt.Println(mapf(nums, tlAdd5))
}

func tlAdd5(n int) int {
	return n + 5
}

func mapf(list []int, tl func(int) int) []int {
	rt := []int{}
	for _, e := range list {
		rt = append(rt, tl(e))
	}
	return rt
}
