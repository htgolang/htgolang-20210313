package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("big.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	/*
		1M:
			$ time go run timeCopyBufferSize.go
			一共读取136.537049秒

			real    2m28.507s
			user    0m0.000s
			sys     0m0.031s

		10M:
			$ time go run timeCopyBufferSize.go
			一共读取116.078072秒

			real    2m0.115s
			user    0m0.000s
			sys     0m0.015s

		100M:
			$ time go run timeCopyBufferSize.go
			一共读取108.501007秒

			real    1m54.923s
			user    0m0.000s
			sys     0m0.015s
	*/

	ctx := make([]byte, 1024*1024)
	for {
		_, err = file.Read(ctx)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
	fmt.Printf("一共读取%f秒\n", time.Now().Sub(start).Seconds())
}
