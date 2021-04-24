package main

import (
	"bytes"
	"fmt"
)

func main() {
	reader := bytes.NewReader([]byte("1234567890"))
	data := make([]byte, 3)

	n, err := reader.Read(data)
	fmt.Println(data, n, err)

	n, err = reader.Read(data)
	fmt.Println(data, n, err)

	n, err = reader.Read(data)
	fmt.Println(data, n, err)

	n, err = reader.Read(data)
	fmt.Println(data, n, err)

	n, err = reader.Read(data)
	fmt.Println(data, n, err)

	reader.Reset([]byte("abc"))

	n, err = reader.Read(data)
	fmt.Println(data, n, err)

	n, err = reader.Read(data)
	fmt.Println(data, n, err)

}
