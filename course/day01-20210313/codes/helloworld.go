/*
声明包, go要求每一个源码都有属于的包
格式: package 包名
包名 目前阶段定义位main
声明包必须在源码中第一行，包声明前可以包含注释
*/
package main

// 后面都是注释 行注释
/*
中间的内容都是注释 块注释
*/

/*
	导入包
	使用标准包或第三方包需要先导入
*/
import (
	"fmt"
)

/*
定义变量(常量)，定义函数，定义类型
*/

/*
定义函数
格式: func 函数名称(参数) {
	函数体
}

目前阶段认为 main函数程序入口
*/
func main() {

	// 调用fmt包中的Println函数
	// 函数功能：在控制台打印出数据(hello world)
	fmt.Println("hello world")
}
