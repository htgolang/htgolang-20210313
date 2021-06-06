package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	bye = "bye"
)

func read(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	txt, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(txt), nil
}

func write(conn net.Conn, txt string) error {
	_, err := fmt.Fprintln(conn, txt)
	return err
}

func input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func handle(conn net.Conn) {
	defer conn.Close()
	fmt.Println("客户端连接:", conn.RemoteAddr())
	for {
		txt, err := read(conn)
		if err != nil {
			fmt.Println("读取客户端数据错误:", err)
			break
		}

		if txt == bye {
			fmt.Println("客户关闭连接")
			break
		}
		fmt.Println("客户端说: ", txt)
		// 服务端回消息
		// 从控制台读取(1行)数据并恢复
		txt = input("回复消息: ")
		// 发送数据到客户端
		err = write(conn, txt)
		if err != nil {
			fmt.Print("发送客户端数据错误:", err)
			break
		}
	}
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
