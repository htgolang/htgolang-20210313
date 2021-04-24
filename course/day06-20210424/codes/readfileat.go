package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path := "readfile.txt"
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	data := make([]byte, 3)

	//var offset int64 = 0
	offset := int64(2)
	n, err := file.ReadAt(data, offset) // 读取err == nil, err != nil && n > 0
	fmt.Println(data, n, err)
}
