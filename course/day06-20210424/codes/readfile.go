package main

import (
	"fmt"
	"io"
	"os"
)

func call() {
	for {
		func() {
			defer func() {
				recover()
			}()
			test()
		}()

	}
}

func test() {
	// open()
	// defer close()
	// readfile // panic()
	//
}

func main() {
	path := "readfile.txt"

	// 相对路径
	// 相对谁的路径
	// 		readfile.go 不可能
	// 		readfile二进制程序(文件路径) 不是
	// 		cd dir; readfile; dir路径(程序运行路径)
	// 绝对路径（固定，程序在哪里运行，程序放在哪里）

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		defer file.Close()
		// file.Close()

		fmt.Println("success")
		// 读文件内容
		data := make([]byte, 3)
		for {
			n, err := file.Read(data)
			if err != nil {
				// 出错
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			} else {
				fmt.Println(data[:n], string(data[:n]))
			}
		}
		// 读取文件内容到data, 内容多余3个字节, 内容少于3个字节，内容3个字节
		// n读取字节的数量
		// error: EOF(文件已经结束)
		// io.EOF
		// fmt.Println(data, n, err)
		// fmt.Println(data[:n])

		// fmt.Printf("%q, %q\n", string(data), string(data[:n]))

		// n, err = file.Read(data)
		// fmt.Println(data, n, err) // data: [49, 0, 0] , n: 0, err: EOF

		// file.Close()
	}
}
