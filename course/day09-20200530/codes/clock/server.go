package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handle(conn net.Conn) {
	defer conn.Close()
	// 处理业务逻辑
	// fmt.Println("local addr::", conn.LocalAddr())
	// fmt.Println("remote addr:", conn.RemoteAddr())
	fmt.Println("client connected:", conn.RemoteAddr())
	fmt.Fprintf(conn, time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
	addr := ":8888"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("listen:", listener.Addr())

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accept: %s", err)
			break
		}
		go handle(conn)
	}
}
