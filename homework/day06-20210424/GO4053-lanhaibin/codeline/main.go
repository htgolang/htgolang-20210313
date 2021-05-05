package main

import (
	"codeline/line"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s /path/to/fileordir\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Println(line.DirLine(os.Args[1]))
}
