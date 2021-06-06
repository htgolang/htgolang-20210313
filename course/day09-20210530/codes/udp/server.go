package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// UDP监听包
	addr := ":9999"
	server, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer server.Close()

	for {
		// 接收数据包（发送者信息addr，发送内容）
		buffer := make([]byte, 1024)
		n, clientAddr, err := server.ReadFrom(buffer)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("recv: ", clientAddr, string(buffer[:n]))
			fmt.Println(server.WriteTo([]byte("ok"), clientAddr))
		}
	}

}
