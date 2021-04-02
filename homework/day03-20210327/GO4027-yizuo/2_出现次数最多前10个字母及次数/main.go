package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

// 查找所有的数值跟出现次数
func findLetterNum(data string) (data_map map[string]int) {
	// 将数据字母转换为小写
	data = strings.ToLower(data)

	// 将文章字符转换为字典并计数
	data_map = map[string]int{}
	for _, data := range data {
		// 去掉特殊字符，只对字母做计数
		if unicode.IsLetter(data) {
			data_map[string(data)]++
		}
	}
	return data_map
}

// 根据传递的切片，针对value int做排序并打印前十的值
func sortMapValue(data map[string]int) {
	// 获取对应的value并赋值给新的int切片
	intSlice := make([]int, 0)
	for _, v := range data {
		intSlice = append(intSlice, v)
	}

	// 切片排序
	intSlice = bubbleSort(intSlice)

	// 顺序打印前十的结果
	for i := 0; i < 10; i++ {
		for k, v := range data {
			if intSlice[i] == v {
				fmt.Println(k, v)
			}
		}
	}

}

// 冒泡排序
func bubbleSort(slice []int) (data []int) {
	var num = len(slice)
	data = make([]int, num)
	copy(data, slice)
	for i := 0; i < num; i++ {
		for k, _ := range data[0:(num - i - 1)] {
			// < 代表倒序；> 代表正序。
			if data[k] < data[k+1] {
				data[k], data[k+1] = data[k+1], data[k]
			}
		}
	}
	return data
}

func main() {
	// 获取文件内容
	data, err := ioutil.ReadFile("I_HAVE_A_DREAM.txt")
	if err != nil {
		errors.New("文件写入异常")
	}

	// 获取所有字符以及对应的出现次数
	dataMap := findLetterNum(string(data))

	// 顺序打印前十的结果
	sortMapValue(dataMap)

}
