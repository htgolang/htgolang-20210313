package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	counts := make(map[string]int)
	re := "abcdefghijklmnopqrstvwuxyzABCDEFGHIUJKLMNOPQRSTUVWXYZ"
	//读取文件内容
	txt, err := ioutil.ReadFile("dream.txt")
	//判断文件是否为空
	if err != nil {
		fmt.Printf("haha %v", err)
		return
	}
	//遍历，并将byte转成字符串类型
	for _, v := range txt {
		s := string(v)
		counts[s] += 1
	}
	//删除非英文单字键
	for k, _ := range counts {
		if  !strings.Contains(re, k) {
			delete(counts, k)
		} 
	}
	fmt.Println(counts)
}