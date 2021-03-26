package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//文件路径
var file_path = "a_dream.txt"

func main() {
	//从文件读取内容
	bytes_content, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	var map1 = map[string]int{}

	for _, v := range bytes_content{
		if (v >= 64 && v <= 90) || (v>=97 &&v<=122){
			map1[string(v)]++
		}
	}

	for k,v := range map1{
		fmt.Printf("字符%s的个数是%d\n", k, v)
	}

}
