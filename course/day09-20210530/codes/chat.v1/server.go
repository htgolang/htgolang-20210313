package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	sizeBufferLength = 5
	bye              = "bye"
)

func read(conn net.Conn) (string, error) {
	sizeBuffer := make([]byte, sizeBufferLength)
	n, err := conn.Read(sizeBuffer)
	if err != nil || n != sizeBufferLength {
		return "", fmt.Errorf("size error")
	}
	size, err := strconv.Atoi(string(sizeBuffer))
	if err != nil {
		return "", err
	}

	contextBuffer := make([]byte, size)
	n, err = conn.Read(contextBuffer)
	if err != nil || n != size {
		return "", fmt.Errorf("content length err")
	}
	return string(contextBuffer), nil
}

func write(conn net.Conn, txt string) error {
	// 先写长度
	size := len(txt) // 5
	// 检查长度
	if float64(size) >= math.Pow10(sizeBufferLength) { // 10 5次方
		// if size > 99999 { // 100000 - 1
		return fmt.Errorf("error write long size")
	}

	formatter := fmt.Sprintf("%%0%dd", sizeBufferLength)
	n, err := fmt.Fprintf(conn, formatter, size)
	if err != nil || n != sizeBufferLength {
		return fmt.Errorf("error write size")
	}

	// 再写内容
	n, err = fmt.Fprintf(conn, txt)
	if err != nil || n != size {
		return fmt.Errorf("error write content")
	}
	return nil
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
