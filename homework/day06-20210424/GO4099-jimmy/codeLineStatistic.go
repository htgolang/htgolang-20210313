package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func lineStatistic(filename string) {
	lines := 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		lines++
	}

	fmt.Printf("文件:%v 行数:%d\n", filename, lines)
}

func dir(dirname string) {
	dirEntrys, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, dirEntry := range dirEntrys {
		if dirEntry.IsDir() {
			dir(filepath.Join(dirname, dirEntry.Name()))
		} else {
			lineStatistic(filepath.Join(dirname, dirEntry.Name()))
		}
	}
}

func main() {
	fmt.Print("请输入目录/文件/文件名后缀: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if strings.Contains(input, "*.") {
		files, err := filepath.Glob(input)
		if err != nil {
			log.Fatal(err)
		}
		if len(files) == 0 {
			log.Fatal("您输入的文件不存在")
		}
		for _, file := range files {
			lineStatistic(file)
		}
	} else {
		fileinfo, err := os.Stat(input)
		if err != nil {
			if os.IsNotExist(err) {
				log.Fatal("您输入的文件不存在")
			}
			log.Fatal(err)
		} else {
			if fileinfo.IsDir() {
				dir(fileinfo.Name())
			} else {
				lineStatistic(fileinfo.Name())
			}
		}
	}
}
