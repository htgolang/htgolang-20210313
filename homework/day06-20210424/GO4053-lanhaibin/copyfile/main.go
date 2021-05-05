package main

import (
	"copyfile/copy"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s src_file dest_file\n", os.Args[0])
	}
	src := os.Args[1]
	dest := os.Args[2]
	copy.Copy(src, dest, 4096)
}
