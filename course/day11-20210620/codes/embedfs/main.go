package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"os"
)

// 多个文件
//go:embed version
//go:embed name
//go:embed *.params
//go:embed params
//go:embed params2/*.txt
var fs embed.FS

func printFile(name string) {
	fmt.Printf("file %s: \n", name)
	file, err := fs.Open(name)
	if err != nil {
		log.Println(err)
		return
	}
	io.Copy(os.Stdout, file)
	fmt.Println()
	file.Close()
}

func main() {
	// printFile("version")
	// printFile("name")
	// printFile("a.params")
	// printFile("b.params")
	// printFile("params/a")
	// printFile("params/b")
	// printFile("params/ctx/a")
	// printFile("params/ctx/b")
	// printFile("params2/a.txt")
	// printFile("params2/a")

	dirs, err := fs.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, dir := range dirs {
		fmt.Println(dir.Name(), dir.IsDir())
	}
}
