package main

import "fmt"


/*
最大元素
 */

func main()  {
	var s []int = []int{9, 8, 19, 100, 2, 8}
	var max,max_index int
	for idx,val := range s {
		if val > max {
			max = val
			max_index = idx
		}
	}
	fmt.Printf("最大数字是%d,下标%d\n",max,max_index)
}
