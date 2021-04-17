package main

import "fmt"

// 别名
type Counter = int

func main() {
	var (
		b byte = 'a'
		r rune = '我'
	)
	fmt.Printf("%T %T\n", b, r)

	var count Counter
	fmt.Printf("%T\n", count)
	fmt.Println(count) // 0
	count++
	fmt.Println(count) // 1
	count += 1
	fmt.Println(count) // 2
	count *= 5
	fmt.Println(count) // 10
}
