package main

import "fmt"

func main() {
	var num int = 64

	num1 := num

	// 修改
	num1 = 165

	// 对num有无影响
	fmt.Println(num1, num) // 无影响

	var pointerNum *int // var pointerNum *int = nil

	fmt.Printf("%T, %v\n", pointerNum, pointerNum)

	pointerNum = &num
	fmt.Printf("%T, %v\n", pointerNum, pointerNum)

	*pointerNum = 116

	fmt.Println(num) // 116 64?

	pointer2Num := &num

	fmt.Printf("%T, %v\n", pointer2Num, pointer2Num)

	*pointer2Num = 1100
	fmt.Println(num)

	fmt.Println(*pointer2Num, *pointerNum)

	var pointer3Num = &num

	fmt.Printf("%T, %v\n", pointer3Num, pointer3Num)

	pp := &pointer2Num
	fmt.Printf("%T, %v\n", pp, pp)
	// 先定义变量 => 取引用(取地址)
	// new(T)
	pointer := new(int)
	fmt.Printf("%T, %v, %v\n", pointer, pointer, *pointer)

	fmt.Printf("%p\n", pointer)

	// usafe
}
