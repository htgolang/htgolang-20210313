package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
)

func copyFile(srcfile, destfile string) {
	src, err := os.Open(srcfile)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	dest, err := os.Create(destfile)
	if err != nil {
		log.Fatal(err)
	}
	defer dest.Close()

	reader := bufio.NewReader(src)
	writer := bufio.NewWriter(dest)
	ctx := make([]byte, 1024*1024)

	for {
		n, err := reader.Read(ctx)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		writer.Write(ctx[:n])
	}
	writer.Flush()
}

func copyDir(src, dest string) {
	files, err := os.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(dest), 0755); err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			copyDir(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))
		} else {
			copyFile(filepath.Join(src, file.Name()), filepath.Join(dest, file.Name()))
		}
	}
}

func main() {
	src := "abc"
	dest := "bak"
	fileInfo, err := os.Stat(src)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("源文件不存在")
		}
		log.Fatal(err)
	} else {
		if fileInfo.IsDir() {
			copyDir(src, dest)
		} else {
			copyFile(src, dest)
		}
	}
}
