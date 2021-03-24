package main

import "fmt"

func main() {
	var data = []int{9, 8, 19, 10, 2, 8}
	for i := 0; i < len(data)-1; i++ {
		var flag = true
		for j := 0; j < len(data)-1-i; j++ {
			//因为这里j+1,防止溢出,所以len(data)-1
			if data[j] > data[j+1] {
				tmp := data[j]
				data[j] = data[j+1]
				data[j+1] = tmp
				flag = false
			}
		}
		//flag为true时,说明已经排好序,直接退出即可
		if flag {
			break
		}
	}
	fmt.Println(data)
}
