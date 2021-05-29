package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// 连接服务器
	addr := "127.0.0.1:9999"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	start := time.Now()
	// 交互
	bytes := make([]byte, 1024)
	n, err := conn.Read(bytes)
	if err == nil {
		fmt.Println(string(bytes[:n]))
	} else {
		fmt.Println(err)
	}
	fmt.Println(time.Now().Sub(start))
	// 关闭连接
	conn.Close()
}
