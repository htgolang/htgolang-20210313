package main

import (
	"io"
	"log"
	"os"
)

func main() {
	//copyfile src dst
	if len(os.Args) != 3 {
		log.Fatal("usage: copyfile srcfile dstfile")
	}

	srcFile, dstFile := os.Args[1], os.Args[2]

	src, err := os.Open(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	bufferSize := 10

	data := make([]byte, bufferSize) // 1M 1024KB => 1024次 2048 => 512次
	// copyfile 10G文件 buffersize 1M, 10M, 100M

	for {
		n, err := src.Read(data)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		n, err = dst.Write(data[:n])
		if err != nil {
			log.Fatal(err)
		}
	}

}
