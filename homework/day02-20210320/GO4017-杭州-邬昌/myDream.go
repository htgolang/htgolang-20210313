package main

import (
	"fmt"
	"io/ioutil"
	"unicode"
)

// 统计一个字符串中每个单词出现的次数。例如："Hello, how do you do?"中，how=1,do=2,you=1

func main() {

	// 通过ioutil 读取本地文件 “dream.txt” 内容，返回一个 byte 切片
	content, err := ioutil.ReadFile("Dream.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}

	// 定义一个map[string]int 进行统计字符
	var map1 =make(map[string]int)

	for _, w := range string(content) { // 遍历字符串
                //通过unicode.IsLetter判断是否是字符
		if unicode.IsLetter(w){

		   //在map中查找，存在，计数加一，不存在，证明刚轮询到一次，计数等于1
		   v,ok:= map1[string(w)]
                   if ok{

			map1[string(w)] = v+1
		   } else {

			map1[string(w)] = 1
			}
		}

	}

	//统计每个字符(大小写英文字符)
	for key, value := range map1 {
		fmt.Printf("%s: %d\n", key, value)
	}
}
