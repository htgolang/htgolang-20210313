package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	start := time.Now()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	resp, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", resp)
	fmt.Println(time.Now().Sub(start))
}
