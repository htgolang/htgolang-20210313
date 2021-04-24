package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, _ := os.Open("bufio.txt")

	// reader := bufio.NewReader(file)
	reader := bufio.NewReaderSize(file, 18)

	// strace

	data := make([]byte, 3)
	n, _ := reader.Read(data)
	fmt.Println(data[:n])

	n, _ = reader.Read(data)
	fmt.Println(data[:n])

	// ctx, isPrefix, _ := reader.ReadLine() // 再缓冲区中数据处理完成后若无换行符就会返回，isPerfix为true，需要继续读取
	// fmt.Println(string(ctx), isPrefix)

	// ctx, isPrefix, _ = reader.ReadLine()
	// fmt.Println(string(ctx), isPrefix)

	// ctx, isPrefix, err := reader.ReadLine()
	// fmt.Println(string(ctx), isPrefix, err)

	// ctx, isPrefix, err = reader.ReadLine()
	// fmt.Println(string(ctx), isPrefix, err)

	// ctx, isPrefix, err = reader.ReadLine()
	// fmt.Println(string(ctx), isPrefix, err) // err EOF ctx无数据

	// fmt.Println(reader.ReadString('\n')) // 最后的换行符保留
	// fmt.Println(reader.ReadString('\n')) // 最后的换行符保留
	// fmt.Println(reader.ReadString('\n')) // 最后的换行符保留
	// fmt.Println(reader.ReadString('\n')) // 最后的换行符保留
	// fmt.Println(reader.ReadString('\n')) // 最后的换行符保留
	// fmt.Println(reader.ReadString('\n')) // 最后的换行符保留

	reader.WriteTo(os.Stdout)

}
