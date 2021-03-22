package main

/*
	create-date:2021-03-21
	describe:统计大小写英文字符数量
	author:GO4035-huangxin
*/

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 通过ioutil 读取本地文件 “dream.txt” n内容，返回一个 byte 切片
	b, err := ioutil.ReadFile("dream.txt")
	if err != nil {
		fmt.Print(err)
	}

	// 定义一个map[string]int 进行统计字符
	tjChar := make(map[string]int, 0)

	for _, v := range b {
		// a-z 的字符编码为 65~90 A-Z 的字符编码为97~122
		if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {

			tjChar[string(v)]++

		}
	}
	fmt.Println("=========统计每个字符(大小写英文字符)=========")
	// 打印统计结果
	for k, v := range tjChar {
		fmt.Printf("字符:%s\t数量:%d\n", k, v)
	}
}
