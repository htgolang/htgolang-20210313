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

func main() {
	addr := "127.0.0.1:8888" // 需要连接的地址
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		txt := input("发送内容:")
		err := write(conn, txt)
		if err != nil {
			fmt.Println("发送消息到服务端失败", err)
			break
		}
		if txt == bye {
			break
		}
		txt, err = read(conn)
		if err != nil {
			fmt.Println("读取消息到服务端失败", err)
			break
		}
		fmt.Println("服务端回复消息:", txt)
	}
}
