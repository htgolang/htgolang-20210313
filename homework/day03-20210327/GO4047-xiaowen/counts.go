package main

import (
	// "bytes"
	"fmt"
	"io/ioutil"

	// "sort"
	"strings"
)


func sorted(nums [][2]int) [][2]int {
	for i := 0;i<len(nums); i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i][1] < nums[j][1]{
				nums[i],nums[j] = nums[j],nums[i]
			}
		}
	}
	return nums
}


func generateMap() [][2]int{
	counts := make(map[string]int)
	cc :=make(map[byte]int)
	var nums [][2]int
	re := "abcdefghijklmnopqrstvwuxyzABCDEFGHIUJKLMNOPQRSTUVWXYZ"
	//读取文件内容
	txt, err := ioutil.ReadFile("dream.txt")
	//判断文件是否为空
	if err != nil {
		fmt.Printf("haha %v", err)
		panic("文件为空") 
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
	
	//将字符串K转换成byte添加到而原数组中
	for k,v := range counts{
		s := []byte(k)
		cc[s[0]] = v
		// fmt.Printf("%v : %v\n", k, v)
	}

	//将byte K转换成int添加到最终数组中
	for k,v := range cc{
		kv := [2]int{int(k), v}
		nums = append(nums, kv)
	}
	return nums
}

func main() {
	
	nums := generateMap()
	nums = sorted(nums)

	//得到top10并打印结果
	for _, v := range nums[:10]{
		fmt.Println(string(byte(v[0])), v[1])
	}
}