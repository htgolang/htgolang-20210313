package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func checkChar(v byte) bool {
	if (v >= 65 && v <= 90) || (v >= 97 && v <= 122) {
		return true
	}
	return false
}

func countMap() map[rune]int {
	file, err := os.Open("dream")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()
	text, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	countMap := make(map[rune]int)
	for _, v := range text {
		if checkChar(v) {
			countMap[rune(v)]++
		}
	}
	return countMap
}

func MySort(arr [][2]int) {
	for i := 0; i < len(arr); i++ {
		tmp := arr[i][1]
		for j := i - 1; j >= 0; j-- {
			if arr[j][1] < tmp {
				arr[j+1][1], arr[j][1] = arr[j][1], tmp
			}
		}
	}
}

func main() {
	var nums [][2]int

	count := countMap()
	for i, v := range count {
		nums = append(nums, [2]int{int(i), v})
	}
	MySort(nums)
	for i, v := range nums {
		if i > 9 {
			break
		}
		fmt.Println(string(v[0]), v[1])
	}
}
