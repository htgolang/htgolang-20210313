package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:8888" // 需要连接的地址
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 处理业务(交换数据)
	buffer := make([]byte, 1023)
	n, err := conn.Read(buffer)
	if err == nil {
		fmt.Println(string(buffer[:n]))
	}

	// 关闭连接
	// conn.Close()
}
