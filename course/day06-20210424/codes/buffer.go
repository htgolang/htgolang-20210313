package main

import (
	"bytes"
	"fmt"
)

func main() {
	buffer := bytes.NewBuffer([]byte("123"))
	data := make([]byte, 2)
	n, err := buffer.Read(data)
	fmt.Println(data, n, err)

	n, err = buffer.Read(data)
	fmt.Println(data, n, err)

	n, err = buffer.Read(data)
	fmt.Println(data, n, err)

	buffer.WriteString("abc123")

	n, err = buffer.Read(data)
	fmt.Println(string(data), n, err)

	n, err = buffer.Read(data)
	fmt.Println(string(data), n, err)

	buffer.WriteString("xyz")

	n, err = buffer.Read(data)
	fmt.Println(string(data), n, err)

}
