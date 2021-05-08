package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=100

func HttpGet(url string) (string, error) {
	var result string
	reps, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer reps.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err := reps.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		result += string(buf[:n])
	}
	return result, nil
}

func Work(start, end int) {
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((start-1)*50)
	for i := start - 1; i < end; i++ {
		ret, err := HttpGet(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		file := fmt.Sprintf("%dpage.html", i)
		writerFile(file, ret)
	}
}

func writerFile(fileName, data string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString(data); err != nil {
		fmt.Println(err)
		return
	}

}

func main() {
	Work(1, 1)
}
