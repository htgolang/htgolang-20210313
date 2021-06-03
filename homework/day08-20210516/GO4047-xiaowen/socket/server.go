package main

import (
	"fmt"
	"log"
	"net"
	"time"
)


func main() {
	listener,err := net.Listen("tcp",":9999")
	if err != nil{
		log.Fatalln(err)
	}

	// b := make([]byte,1024)
	// _, err = conn.Read(b)
	// fmt.Println(err)
	// if err !=nil{
	// 	log.Fatalln(err)
	// }
	// fmt.Println(string(b))
	for {
		conn,err := listener.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}

		go func ()  {
			//打印客户端ip
			time.Sleep(10 * time.Second)
			log.Println("client:", conn.RemoteAddr())
			//向客户端发送数据
			fmt.Fprintf(conn, time.Now().Format("2006-01-02 15:04:05"))
			//关闭连接
			conn.Close()
		}()

	}
	
	listener.Close()
}
