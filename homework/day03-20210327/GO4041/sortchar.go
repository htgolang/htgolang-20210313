package main

import (
	"fmt"
	"io/ioutil"
)

// getchamap 读取文件，输出map[int]int
func getchamap(fs string) map[int]int {
	// 申明map
	rep := make(map[int]int)

	// 打开读取文件
	content, err := ioutil.ReadFile(fs)
	if err != nil {
		panic(err)
	}

	// 遍历统计
	for _, i := range string(content) {
		if (i >= 65 && i <= 90) || (i >= 97 && i <= 122) {
			v, ok := rep[int(i)]
			if ok {
				rep[int(i)] = v + 1
			} else {
				rep[int(i)] = 1
			}
		}
	}

	// 返回
	return rep
}

// sortcha 读取map[int]int 输出排序[][2]int
func sortcha(nums map[int]int) {
	var rep [][2]int
	// 初始化[][2]int
	for k, v := range nums {
		tn := [2]int{k, v}
		rep = append(rep, tn)
	}
	// 插入排序 输出[][2]int
	for k, v := range rep {
		if k > 0 {
			for i := k - 1; i >= 0; i-- {
				if rep[i][1] > v[1] {
					rep[i][1], rep[i+1][1] = v[1], rep[i][1]
				}
			}
		}
	}

	// 输出后面10个索引的值,并且转换为str输出
	for i := len(rep) - 1; i > len(rep)-11; i-- {
		fmt.Println(string(rep[i][0]), rep[i][1])
	}
	// fmt.Println(rep)

}

func main() {
	m := getchamap("i-have-a-dream.txt")
	sortcha(m)

}
