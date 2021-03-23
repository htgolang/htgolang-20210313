package main

import "fmt"

func main() {
	queue := []string{}

	var txt string
	/*
		输入do执行任务，从队列中拿去一个
		exit退出
		其他添加任务
	*/
	for {
		fmt.Print("请输入任务(exit退出,do执行):")
		fmt.Scan(&txt)
		if txt == "exit" {
			break
		} else if txt == "do" {
			if len(queue) == 0 {
				fmt.Println("无任务")
			} else {
				fmt.Println("执行:", queue[0])
				queue = queue[1:]
			}
		} else {
			queue = append(queue, txt)
		}
	}

}
