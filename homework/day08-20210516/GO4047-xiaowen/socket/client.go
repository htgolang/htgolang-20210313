package main

import (
	"fmt"
	"log"
	"net"
)


func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil{
		log.Fatalln(err)
	}
	// _,err = conn.Write([]byte("hahah"))
	// if err != nil{
	// 	log.Fatalln(err)
	// }
	b := make([]byte, 1024)
	_,err = conn.Read(b)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(string(b))
	defer conn.Close()
}