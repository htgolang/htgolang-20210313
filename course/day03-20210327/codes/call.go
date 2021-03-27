package main

import "fmt"

func test(n int, s []int, p *int) {
	fmt.Printf("test: %p %p %p %p %v\n", &n, &s, s, &p, p)
	n = 2
	s = []int{2, 3, 4}
	p = &n
}

func main() {
	n := 1
	s := []int{1, 2, 3}
	p := &n
	// s => %p 引用底层数组的地址
	fmt.Printf("main: %p %p %p %p %v\n", &n, &s, s, &p, p)
	test(n, s, p)
}
