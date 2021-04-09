package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

			//在map中查找，存在，计数加一，不存在，证明刚轮询到第一次，计数等于1
			v,ok:= map1[string(w)]
			if ok{

				map1[string(w)] = v+1
			} else {

				map1[string(w)] = 1
			}
		}

	}

	//定义一个切片，将map1的value值添加到slice中
	var nums []int
	for _,v :=range map1{

		nums =append(nums,v)
	}
	//对slice进行排序
	sort.Ints(nums)

	//取出排名前十的数字，和map1值进行对比，相等则输出
	for _, value := range nums[len(nums)-10:len(nums)]{
		for k,v :=range map1{
			if v ==value{
				fmt.Printf("%s: %d\n", k, v)
			}

		}
	}

}