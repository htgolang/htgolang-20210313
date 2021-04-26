package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func copyfile(srcFile, destFile string) {
	src, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dest, err := os.Create(destFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	ctx := make([]byte, 1024)
	for {
		n, err := src.Read(ctx)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		dest.Write(ctx[:n])
	}
}

func main() {
	start := time.Now()
	copyfile("big.txt", "big.txt_bak")
	fmt.Printf("复制用时%v\n", time.Now().Sub(start))
}
//1K     复制用时59.169841s     复制用时57.0729823s
//1M     复制用时25.5094037s    复制用时34.3928233s
//10M    复制用时55.2387665s    复制用时43.0317052s
//100M   复制用时32.503246s     复制用时32.9200517s
//1G     复制用时35.4535703s    复制用时34.5435553s
