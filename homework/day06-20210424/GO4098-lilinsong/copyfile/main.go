package main

import (
	"copyfile/utils"
	"flag"
	"fmt"
	"os"
)

func main() {
	var src string
	var dest string
	var cp utils.CopyFile
	flag.StringVar(&src, "src", "", "specify source file location")
	flag.StringVar(&dest, "dest", "", "specify dest file location")

	flag.Usage = func() {
		fmt.Printf("%v --src /path/to --dest /path/to", os.Args[0])
	}
	flag.Parse()

	if src == "" || dest == "" {
		fmt.Println("src 或 dest 不能为空")
		return
	}
	cp.Start(src, dest)
}
