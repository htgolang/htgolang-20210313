package main

import (
	"fmt"
	"os"
)

func main() {
	data := make([]byte, 3)
	n, err := os.Stdin.Read(data)
	fmt.Println(n, err, data)

	os.Stdout.WriteString("你输入的内容为: " + string(data[:n]))

	fmt.Fprintf(os.Stdout, "你输入的内容为: %s", string(data[:n]))
	os.Stderr.WriteString("错误")
}
