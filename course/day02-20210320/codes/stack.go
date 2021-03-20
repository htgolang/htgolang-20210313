package main

import "fmt"

func main() {
	stack := []string{}

	var txt string
	/*
		输入do执行任务，从栈中拿去一个
		exit退出
		其他添加任务
	*/
	for {
		fmt.Print("请输入任务(exit退出,do执行):")
		fmt.Scan(&txt)
		if txt == "exit" {
			break
		} else if txt == "do" {
			if len(stack) == 0 {
				fmt.Println("无任务")
			} else {
				fmt.Println("执行任务:", stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, txt)
		}
	}

}
