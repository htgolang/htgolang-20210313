package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("1234567890")

	data := make([]byte, 3)
	n, err := reader.Read(data)
	fmt.Println(data, err, n)

	n, err = reader.Read(data)
	fmt.Println(data, err, n)

	n, err = reader.Read(data)
	fmt.Println(data, err, n)

	n, err = reader.Read(data)
	fmt.Println(data, err, n)

	n, err = reader.Read(data)
	fmt.Println(data, err, n)

	reader.Seek(0, 0)

	n, err = reader.Read(data)
	fmt.Println(data, err, n)

}
