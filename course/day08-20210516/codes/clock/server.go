package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// 监听IP和端口
	addr := ":9999" //"0.0.0.0:9999"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// 获取客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 启动例程处理客户端连接
		go func() {
			time.Sleep(10 * time.Second)
			log.Println("client:", conn.RemoteAddr())
			// 交互处理
			fmt.Fprintf(conn, time.Now().Format("2006-01-02 15:04:05"))
			// 关闭连接
			conn.Close()
		}()

	}

	// 关闭服务器
	listener.Close()

}
