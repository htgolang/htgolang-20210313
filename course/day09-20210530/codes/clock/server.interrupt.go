package main

import (
	"log"
	"net"
)

func main() {
	addr := "127.0.0.1:8888"

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	interrupt := make(chan int, 1)
	go func() {
		// interrupt 写入终止信号
	}()
END:
	for {
		select {
		case <-interrupt:
			break END
		default:

		}
	}

	listener.Close()
}
