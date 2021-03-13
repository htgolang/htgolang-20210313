package main

import "fmt"

func main() {
	fmt.Println("a", "b")
	fmt.Println("1", "2")
	fmt.Print("a", "b")
	fmt.Print("1", "2")

	// 格式化的字符串 占位, 后面的数据进行填充
	// 占位符
	// %T %v %#v 于数据类型无关
	// 打印变量类型， %v, %#v 以可识别的格式打印
	// %s 字符串
	// %d 整数
	fmt.Println()
	var name = "kk"
	var id = 1
	fmt.Printf("name: %T, id: %T\n", name, id)
	fmt.Printf("name: %v, id: %v\n", name, id)
	fmt.Printf("name: %#v, id: %#v\n", name, id)
	fmt.Printf("name: %s, id: %d\n", name, id)
	fmt.Printf("name: %d, id: %s\n", name, id)
}
