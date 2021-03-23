package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

func find_letter_num(data string) {
	// 将数据字母转换为小写
	data = strings.ToLower(data)

	// 将文章字符转换为字典并计数
	data_map := map[string]int{}
	for _, data := range data {
		// 去掉特殊字符，只对字母做计数
		if unicode.IsLetter(data) {
			data_map[string(data)]++
		}
	}
	fmt.Println(data_map)
}

func main() {
	// 获取文件内容
	data, err := ioutil.ReadFile("I_HAVE_A_DREAM.txt")
	if err != nil {
		fmt.Print(err)
	}

	find_letter_num(string(data))
}
