package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var count int

func codeNum(path string) int {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	if fileInfo.IsDir() {
		fileInfos, _ := os.ReadDir(path)
		for _, file := range fileInfos {
			if file.IsDir() {
				codeNum(filepath.Join(path, file.Name()))
			} else {
				filePath := filepath.Join(path, file.Name())
				if strings.HasSuffix(filePath, ".go") {
					fmt.Println(filePath)
					num, err := readFile(filePath)
					if err != nil {
						fmt.Println(err)
					}
					count += num
				}
			}
		}
	} else {
		num, err := readFile(path)
		if err != nil {
			fmt.Println(err)
		}
		count += num
	}
	return count
}

func readFile(path string) (int, error) {
	var num int
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return 0, err
	}
	reader := bufio.NewReader(file)
	for {
		if line, _, err := reader.ReadLine(); err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		} else {
			if string(line) == "\n" {
				continue
			} else {
				num++
			}
		}
	}
	return num, nil
}

func main() {
	ret := codeNum("../password")
	fmt.Println(ret)
}
