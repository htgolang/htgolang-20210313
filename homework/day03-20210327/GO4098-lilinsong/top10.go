package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func traverseArticle(path string) [][2]int {
	var statistics = make(map[int]int, 0)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	txt, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(txt); i++ {
		if (txt[i] >= 65 && txt[i] <= 90) || (txt[i] >= 97 && txt[i] <= 122) {
			statistics[int(txt[i])]++
		}
	}
	ret := top10(statistics)
	return ret[:10]
}

func top10(dict map[int]int) [][2]int {
	arr := make([][2]int, len(dict))
	var index int
	for k, v := range dict {
		arr[index] = [2]int{k, v}
		index++
	}

	return quickSort(arr, 0, len(arr)-1)

}

func quickSort(arr [][2]int, left, right int) [][2]int {
	if left < right {
		mid := partitions(arr, left, right)
		quickSort(arr, left, mid-1)
		quickSort(arr, mid+1, right)
	}
	return arr
}

func partitions(arr [][2]int, left, right int) int {
	tmp := arr[left]
	for left < right {
		for left < right && arr[right][1] <= tmp[1] {
			right--
		}
		arr[left] = arr[right]

		for left < right && arr[left][1] >= tmp[1] {
			left++
		}
		arr[right] = arr[left]
	}
	arr[left] = tmp

	return left
}

func main() {
	path := "G:/go-work/src/github.com/test/test02/dream.txt"
	ret := traverseArticle(path)
	num := 1
	for _, v := range ret {
		fmt.Printf("top%d:\t%s 出现 %d\n", num, string(rune(v[0])), v[1])
		num++
	}
}
