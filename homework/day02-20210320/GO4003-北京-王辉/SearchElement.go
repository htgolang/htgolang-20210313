package main

import "fmt"

/*
查找元素:
这个有问题，没做出来
 */
var ret int
func main() {
	s := []int{2, 8, 9, 10, 19}
	input := 0
	fmt.Scanf(">>查找元素的位置",&input)
	var ret int
	for idx,val := range s {
		if val == input{
			ret = idx
		}else {
			ret = -1
		}
	}
	fmt.Printf("%d的索引是%d\n",input,ret)
}
