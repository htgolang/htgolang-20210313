package main

import "fmt"
/*
倒数第二大换到倒数第二位
 */
func main() {
	var max,second,second_index int
	var s []int = []int{9, 8, 19, 10, 2, 8}
	fmt.Printf("调整前的值%#v\n",s)
	for idx,val := range s {
		if val > max {
			max = val
			second =val
			//second_index = idx
		}else {
			if second < max {
				second = val
				second_index = idx
			}
		}
	}
	s[second_index],s[len(s)-2] = s[len(s)-2],s[second_index]
	fmt.Printf("第二大数字是%d,索引是%d\n",second,second_index)
	fmt.Printf("调整后的值%#v\n",s)
}