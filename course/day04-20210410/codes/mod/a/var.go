package a

import "fmt"

const Name = "a"

func Call() {
	fmt.Println("a")
}

// 需要调用一些函数进行计算进行初始化
var Version string

// go 定义的
func init() {
	fmt.Println("a. init")
	Version = "1.10.0"
}
