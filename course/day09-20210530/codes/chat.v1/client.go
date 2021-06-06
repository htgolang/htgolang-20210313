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
