package main

import (
	"fmt"
	"io/ioutil"
)

func bubbleSort(p [][2]int) [][2]int {
	for i := 0; i < len(p)-1; i++ {
		for j := 1 + i; j < len(p); j++ {
			if p[i][1] < p[j][1] {
				p[i], p[j] = p[j], p[i]
			}
		}
	}
	return p
}

func main() {
	data, err := ioutil.ReadFile("dream.txt")
	if err != nil {
		panic(err)
	}

	s := make(map[byte]int)
	for _, v := range data {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' {
			s[v]++
		}
	}

	p := [][2]int{}
	for k, v := range s {
		p = append(p, [2]int{int(k), v})
	}

	p = bubbleSort(p)
	for _, v := range p[:10] {
		fmt.Printf("%c %d\n", v[0], v[1])
	}
}
