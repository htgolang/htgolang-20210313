package main

import (
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("usage: copyfile srcfile dstfile buffersize(int)")
	}
	srcFile, dstFile, _bufferSize := os.Args[1], os.Args[2], os.Args[3]
	//(1)打开源文件
	src, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()
	//(2)创建目的文件
	dst, err := os.Create(dstFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()
	bufferSize, _ := strconv.Atoi(_bufferSize)
	data := make([]byte, bufferSize) // 1M 1024KB => 1024次 2048 => 512次
	// copyfile 10G文件 buffersize 1M, 10M, 100M
	for {
		//从源文件读取内容
		n, err := src.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		//往目的文件写入内容
		n, err = dst.Write(data[:n])
		if err != nil {
			log.Fatal(err)
		}
	}
} /*
dd if=/dev/zero bs=8192 count=1250 of=test1.file
(1)10bytes
❯ time go run test1Going.go test1.file test2.file 10
1.13s user
3.70s system
74% cpu
6.494 total
(2)1024bytes=1k
❯ time go run test1Going.go test1.file test3.file 1024
0.24s user
0.32s system
58% cpu
0.966 total
(3)1024*1024=1048576=1M
❯ time go run test1Going.go test1.file test3.file 1048576
0.23s user
0.27s system
37% cpu
1.364 total
(4)1024*1024*10=10485760=10M
❯ time go run test1Going.go test1.file test4.file 10485760
0.26s user
0.27s system
53% cpu
1.003 total
(5)1024*1024*100=104857600=100M
❯ time go run test1Going.go test1.file test5.file 104857600
0.24s user
0.27s system
15% cpu
3.326 total
*/
