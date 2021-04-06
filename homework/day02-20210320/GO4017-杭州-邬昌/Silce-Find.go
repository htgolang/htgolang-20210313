package main

import "fmt"

func main() {

	silce1 := []int{2,8,9,10,19}
	start,end :=0,len(silce1)


	var num int
	fmt.Println("请输入要找的数字:")
	fmt.Scan(&num)

	for {

		middle := (start + end) / 2

		if num > silce1[middle] {
			start = middle + 1

		} else if num < silce1[middle] {
			end = middle - 1
		} else {
			fmt.Println("找到了,索引为",middle)
			break
		}

		if start>end{
			fmt.Println("未找到,数字不存在")
			break
		}
	}
}


