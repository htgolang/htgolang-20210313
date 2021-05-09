package main

import "fmt"

type EmptyStruct struct{}

type EmptyInterface interface{}

func main() {
	emptyString := ""
	fmt.Println(emptyString)

	emptySlice := make([]int, 0)
	fmt.Println(emptySlice)

	emptyMap := make(map[string]string)
	fmt.Println(emptyMap)

	emptyStruct := EmptyStruct{}
	fmt.Println(emptyStruct)

	var emptyInterface EmptyInterface

	fmt.Println(emptyInterface)

	// 空接口 如何赋值 任意对象都可以赋值给空接口
	emptyInterface = 1
	fmt.Printf("%T, %#v\n", emptyInterface, emptyInterface)
	emptyInterface = "123"
	fmt.Printf("%T, %#v\n", emptyInterface, emptyInterface)
	emptyInterface = true
	fmt.Printf("%T, %#v\n", emptyInterface, emptyInterface)
	emptyInterface = emptyStruct
	fmt.Printf("%T, %#v\n", emptyInterface, emptyInterface)
	emptyInterface = emptySlice
	fmt.Printf("%T, %#v\n", emptyInterface, emptyInterface)

	// 空接口 + 匿名
	var empty interface{} // empty空接口的变量

	fmt.Printf("%T, %#v\n", empty, empty)
	empty = 1
	fmt.Printf("%T, %#v\n", empty, empty)
	empty = true
	fmt.Printf("%T, %#v\n", empty, empty)

	//如何实现一个可以接收任意类型变量的函数
	fmt.Println(t(1))
	fmt.Println(t(""))
	fmt.Println(t(true))
	fmt.Println(t([]string{"1", "@"}))
	fmt.Println(t(map[string]int{"1": 1, "@": 2}))

	// 如何实现一个可以接收任意数量的任意类型的函数
	args(1, 2, 3)
	args("abc", 2, 3, []string{}, struct{}{})
}

func t(arg interface{}) string {
	return fmt.Sprintf("%T", arg)
}

func args(args ...interface{}) {
	for i, arg := range args {
		fmt.Println(i, t(arg))
	}
}
