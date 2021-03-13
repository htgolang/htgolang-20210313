package main

import "fmt"

func main() {
	// 老婆
	// 买10个包子, 路上碰到卖西瓜的买一个西瓜
	var confirm string

	fmt.Println("老婆的想法:")
	// 标准输出
	fmt.Println("买10个包子")
	// 从控制台输入输入为:y =>碰到卖西瓜的了
	fmt.Print("有卖西瓜的吗(y/n):")
	fmt.Scan(&confirm)

	// 需要根据confirm与y比较结果判断是否执行
	// confirm == "y" 执行
	/*
		if expr {

		}
	*/
	if confirm == "y" {
		fmt.Println("买1个西瓜")
	}

	// 老公
	// 有卖西瓜的吗， 有 买1个包子，否则买10个包子

	fmt.Println("老公想法:")
	fmt.Print("有卖西瓜的吗(y/n):")
	fmt.Scan(&confirm)

	/*
		if expr {

		} else {

		}
		expr true if语句块，否则执行else语句块
	*/
	// confirm == "y"
	if confirm == "y" {
		fmt.Println("买1个包子")
	} else {
		fmt.Println("买10个包子")
	}

}
