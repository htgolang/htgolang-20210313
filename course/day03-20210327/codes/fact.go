package main

import "fmt"

func fact(n int64) int64 {
	fmt.Println("call", n)
	if n == 0 || n == 1 {
		return 1
	}
	rt := n * fact(n-1)
	fmt.Println("result", n, rt)
	return rt
}

func main() {
	fmt.Println(fact(5))
}
