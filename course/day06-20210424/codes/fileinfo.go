package main

import (
	"fmt"
	"log"
	"os"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

func main() {
	path := "mgr1"
	file, err := os.Stat(path) // os.Lstat =>连接 =>连接文件 Stat => 原文件
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
			return

		}
		log.Fatal(err)
	}
	fmt.Println(file.IsDir())
	
}
