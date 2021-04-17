package main

import "fmt"

// 重定义
type Counter int

// go http
// headers key: value map
// type Header map[string][]string

func main() {
	var i int = 10
	var count Counter

	fmt.Printf("%T, %v\n", count, count) // Counter

	fmt.Println(count)
	count += 1
	fmt.Println(count)
	count *= 2
	fmt.Println(count)

	count += Counter(i) // 左右操作数类型必须一致
	fmt.Println(count)
}
