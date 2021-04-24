package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path := "rdwr.txt"
	flags := os.O_RDWR | os.O_CREATE
	file, err := os.OpenFile(path, flags, os.ModePerm) //文件不存在创建还是报错?
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("success")

	fmt.Println(file.Seek(0, 1))

	file.WriteString("abc")

	fmt.Println(file.Seek(0, 1))

	file.Seek(1, 0) // 文件开始

	data := make([]byte, 3)
	n, err := file.Read(data)
	fmt.Println(data, n, err, string(data[:n]))

	file.WriteString("123")

	file.Seek(1, 0)
	file.WriteString("xyz") //axyz23
}
