package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("usage: tailf web.log")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// 1. 文件被人删除/重建了怎么办 每天一个日志文件 web.log
	// 检查文件是否被删除/重建 => inode
	// 2. 文件被人清空的怎么办
	// 如何检查文件被清空了 size()
	// 3. web.log.xxxx
	//
	reader := bufio.NewReaderSize(file, 16)
	line := make([]byte, 0, 4096)
	for {
		data, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Print(err)
				break
			}
			// break 不退出
			// 程序等待n秒
			time.Sleep(time.Second)
			// } else {
			// 	fmt.Println(string(data))
			// }
		} else if isPrefix == true {
			line = append(line, data...)
		} else {
			if len(line) > 0 {
				fmt.Println(string(append(line, data...)))
				line = make([]byte, 0, 4096)
			} else {
				fmt.Println(string(data))
			}
		}
	}
}
