package main

import "fmt"

/*
第二大元素
*/
func main() {
	var s []int = []int{9, 8, 19, 10, 2, 8}
	max := 0
	second := 0
	second_index := 0

	for i := 0;i <len(s) ;i++  {
		if max < s[i]{
			second = max // 获取上一个最大值
			max = s[i]
		}else {
			// 第二大的数据和后面的数据做对比
			if second < max {
				if second < s[i] {
					second = s[i]
					second_index = i
				}
			}
		}
	}
	fmt.Printf("第二大数字是%d,索引是%d\n",second_index,second)
}
