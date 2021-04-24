package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// var txt string
	// fmt.Scan(&txt)
	// fmt.Println("scan:", txt)
	// 带缓冲区IO
	// 读取字符串数据 strconv转换

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println("输入内容:", scanner.Text(), scanner.Bytes())
	}
	fmt.Println(scanner.Err())

}
