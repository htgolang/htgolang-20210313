package main

import "fmt"

// 参数

// 无参函数
func sayHello() {
	fmt.Println("Hello World")
}

// 有参函数

func sayHi(name string) {
	fmt.Println("Hi", name)
}

func add(left int, right int) {
	fmt.Println("left=", left, "right=", right)
	fmt.Println(left + right)
}

func add3(n1 int, n2 int, n3 int) {
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
	fmt.Println(n1 + n2 + n3)
}

func argstype(a1 bool, a2 string, a3 int, a4 [5]int, a5 []string, a6 map[string]string) {
	fmt.Printf("%T, %T, %T, %T, %T, %T\n", a1, a2, a3, a4, a5, a6)
	fmt.Println(a1, a2, a3, a4, a5, a6)
}

// 多个连续参数类型一致，只保留连续参数中最后一个参数类型，前面n个参数类型省略
/*
var (
	a1 bool
	a2, a3, a4 string
	a5 []int
	a6 map[string]string
)
*/
func merge(a1 bool, a2, a3 string, a4, a5 bool, a6 []int) {
	fmt.Printf("%T %T %T %T %T %T\n", a1, a2, a3, a4, a5, a6)
	fmt.Println(a1, a2, a3, a4, a5, a6)
}

// 可变参数
// 变量名 ...T
// 可以接收任意多个string类型的变量
// 可变参数在一个函数内只能定义一次
func anyArgs(args ...string) {
	// args ?? 什么类型的[]T
	fmt.Printf("%T %v\n", args, args)
}

func argsSlice(args []string) {
	fmt.Printf("%T %v\n", args, args)
}

func addn(args ...int) {
	rt := 0
	for _, arg := range args {
		rt += arg
	}
	fmt.Println("add:", rt)
}

/* 不能定义多个可变参数
func any(argsint ...int, argsstring ...string) {

}
*/

// 可变参数必须放在最末尾
func add2(left int, right int, args ...int) {
	rt := left + right
	for _, arg := range args {
		rt += arg
	}
	fmt.Println("add2:", rt)
}

func main() {
	sayHello()
	sayHi("kk")
	add(1, 2)
	add3(1, 2, 3)
	argstype(false, "xxx", 10, [5]int{}, nil, nil)
	merge(true, "x", "y", true, false, nil)

	anyArgs()
	anyArgs("1")
	anyArgs("1", "2")
	anyArgs("1", "2", "3")
	anyArgs("1", "2", "3", "4")
	sArgs := []string{"a", "b", "c"}

	anyArgs(sArgs...)
	argsSlice(sArgs)

	addn()
	addn(2, 3)
	addn(2, 3, 100, 200)

	add2(1, 2)
	add2(1, 2, 100, 1000)
}
