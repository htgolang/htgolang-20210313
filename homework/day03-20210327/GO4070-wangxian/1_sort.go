package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

//统计文件中每个字母数量
func chatCount(file_path string) map[byte]int {
	//从文件读取内容
	bytes_content, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var map1 = map[byte]int{}

	for _, v := range bytes_content {
		if (v >= 64 && v <= 90) || (v >= 97 && v <= 122) {
			map1[v]++
		}
	}
	return map1
}

func main() {
	var s1 [][2]int
	file_path := "a_dream.txt"
	count_map := chatCount(file_path)

	for k, v := range count_map {
		s1 = append(s1, [2]int{int(k), v})
	}

	sort.Slice(s1, func(i, j int) bool { return s1[i][1] > s1[j][1] })

	s2 := s1[:10]

	fmt.Println("出现最多的10个字母如下:")
	for _, v := range s2 {
		fmt.Printf("%c出现%d次\n", v[0], v[1])
	}

}
