package main

import "fmt"

func main() {
	// 初始化切片
	var s []int = []int{10, 9, 0, 5, 7, 2, 3}

	for k, v := range s {
		if k > 0 {
			// 倒序循环
			// 如果当前值大则不动索引
			// 如果当前值小则交换
			for i := k - 1; i >= 0; i-- {
				if s[i] > v {
					s[i], s[i+1] = v, s[i]
				}
			}
		}
	}
	fmt.Println(s)
}
