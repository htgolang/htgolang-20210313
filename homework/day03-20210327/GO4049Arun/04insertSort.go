package main

import "fmt"

func main() {
	var arr = []int{12, 67, 8, 36, 27, 88, 27, 99, 28}
	var newarr = make([]int,len(arr))
	newarr[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		var cur = arr[i];
		for j := len(newarr) - 1; j >= 0; j-- {
			if (cur > newarr[j]) {
				if (j == 0) {
					newarr = append([]int{cur}, newarr...)
				}
			} else {
				// 在所以j的后面把这一项放进去;
				newarr = append(newarr[:j+1],append([]int{cur},newarr[j+1:]...)...)
				// 一旦给这一项找到了位置,放到了数组中,那么直接停止该小循环;
				break;
			}
		}
	}
	fmt.Println(newarr);//[99 88 67 36 28 27 27 12 8 0 0 0 0 0 0 0 0]
}
