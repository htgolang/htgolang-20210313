package main

import "fmt"

// 生成一个新的函数 每次调用生成一个ID => 从start开始，id的间隔是step
// func() int
// 面向切面编程 装饰器
func generateId(start int, step int) func() int {
	return func() int {
		id := start
		start += step
		return id
	}
}

func main() {
	// 0, 1
	g1 := generateId(0, 1)

	fmt.Println(g1())
	fmt.Println(g1())
	fmt.Println(g1())
	fmt.Println(g1())
	fmt.Println(g1())
	// 100, 100
	g100 := generateId(100, 100)
	fmt.Println(g100())
	fmt.Println(g100())
	fmt.Println(g100())
	fmt.Println(g100())
	fmt.Println(g100())
	fmt.Println(g100())
}
