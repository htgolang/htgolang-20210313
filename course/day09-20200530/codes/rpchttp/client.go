package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 请求对象
type Request struct {
	Left  int
	Right int
}

// 响应对象
type Response struct {
	R1 int
	R2 int
	R3 int
	R4 int
}

func main() {
	addr := "127.0.0.1:9999"
	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	req := Request{10, 2}
	resp := Response{}
	err = client.Call("Calc.Calc", req, &resp)
	if err == nil {
		fmt.Println(resp)
	} else {
		fmt.Println(err)
	}

}
