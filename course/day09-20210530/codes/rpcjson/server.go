package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 计算两个int类型数字的加减乘除

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

// 计算对象
type Calc struct {
}

func (c *Calc) Calc(req Request, resp *Response) error {
	fmt.Println("cacl", req)
	if req.Right == 0 {
		return fmt.Errorf("divide zero")
	}
	resp.R1 = req.Left + req.Right
	resp.R2 = req.Left - req.Right
	resp.R3 = req.Left * req.Right
	resp.R4 = req.Left / req.Right
	return nil
}

func main() {
	addr := ":9999"
	// json rpc
	rpc.Register(&Calc{})

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		go jsonrpc.ServeConn(conn)
	}
}
