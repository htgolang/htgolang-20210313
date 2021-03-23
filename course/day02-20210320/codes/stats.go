package main

import "fmt"

func main() {
	stats := make(map[string]int)
	// 统计所有人上课次数
	// 从控制台上输入一个人名字:
	// exit 结束，并打印每个人上课次数
	// 非exit： 给学生上课次数+1

	var txt string
	for {
		fmt.Print("请输入名称(exit退出):")
		fmt.Scan(&txt)
		if txt == "exit" {
			break
		}
		// 统计
		// stats[txt] = stats[txt] + 1
		// stats[txt] += 1
		stats[txt]++
		// a = a + b => a += b
		fmt.Println(stats)
	}

	for k, v := range stats {
		fmt.Println(k, v)
	}
}
