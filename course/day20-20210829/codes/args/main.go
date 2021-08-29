package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(currentDir)

}
