package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func Line(file string) (lines int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewScanner(f)
	for reader.Scan() {
		lines++
	}
	return lines
}

func DirLine(dir string) (ret int) {
	fslist, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, fs := range fslist {
		if fs.IsDir() {
			ret += DirLine(filepath.Join(dir, fs.Name()))
		} else {
			ret += Line(filepath.Join(dir, fs.Name()))
		}
	}
	return ret
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s /path/to/fileordir\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Println(DirLine(os.Args[1]))
}
