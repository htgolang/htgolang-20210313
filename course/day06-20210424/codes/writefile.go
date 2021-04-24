package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path := "writefile.txt"

	file, err := os.Create(path) // 文件存在 清空，不存在 创建
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	n, err := file.Write([]byte("12345"))
	fmt.Println(n, err)

	n, err = file.Write([]byte("xyz"))
	fmt.Println(n, err)

	n, err = file.WriteString("abcdef")
	fmt.Println(n, err)

	n, err = file.WriteString("我")
	fmt.Println(n, err)

	n, err = file.WriteAt([]byte("???"), 0)
	fmt.Println(n, err)

	data := make([]byte, 10)
	n, err = file.Read(data)
	fmt.Println(data, n, err)
}
