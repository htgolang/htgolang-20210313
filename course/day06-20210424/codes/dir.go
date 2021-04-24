package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open(".")
	if err != nil {
		log.Fatal(file)
	}
	defer file.Close()
	//	读取目录下的子文件(有可能是一个文件，有可能是一个子目录)
	// fmt.Println(file.Readdir(10))
	// fmt.Println(file.Readdirnames(10))
	// fmt.Println(file.Readdirnames(10))
	// fmt.Println(file.Readdirnames(-1))
	files, err := file.Readdir(-1)
	for _, file := range files {
		fmt.Println(file.IsDir(), file.Name(), file.Size(), file.ModTime(), file.Mode())
	}
}
