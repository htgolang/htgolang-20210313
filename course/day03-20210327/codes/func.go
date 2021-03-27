package main

import "fmt"

// 调用时打印helloworld字符串在控制台
// sayHello
// 无参 无返回值
func sayHello() {
	fmt.Println("Hello World")
}

// 调用时传递用户名，在控制台打印HI xxx
// sayHi
// 参数 定义在空号内 名称 类型
func sayHi(user string) { // 形参
	fmt.Println("Hi", user)
}

// 有参 有返回值
// add函数 计算两个int类型的和 并返回
func add(left int, right int) int {
	return left + right
}

func main() {
	sayHello()
	// 从控制台输入一个用户名
	// 调用函数打印 hi xxx
	var txt string
	fmt.Print("请输入用户名：")
	fmt.Scan(&txt)
	sayHi(txt) // 实参 => 变量赋值

	rt := add(5, 10)
	fmt.Printf("%T\n", rt)
}
