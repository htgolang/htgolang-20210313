package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:9999")
	if err != nil {
		log.Fatal(conn)
	}
	defer conn.Close()

	fmt.Println(conn.Write([]byte(time.Now().Format("2006-01-02 15:04:05"))))
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	fmt.Println(string(buffer[:n]))
}
