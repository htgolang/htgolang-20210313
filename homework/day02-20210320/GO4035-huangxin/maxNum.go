package main

/*
	create-date:2021-03-21
	describe:找最大值与其索引
	author:GO4035-huangxin
*/

import "fmt"

func main() {
	/*
				3. 找到切片中最大的元素索引和值(打印)
		    	[]int{9, 8, 19, 10, 2, 8}
	*/
	nums := []int{9, 8, 19, 10, 2, 8}

	max_num := 0
	max_index := 0
	for k, v := range nums {
		if max_num < v {
			max_num = v
			max_index = k
		}
	}

	fmt.Printf("nums 中最大值为 %d, 其索引为 %d\n", max_num, max_index)
}
