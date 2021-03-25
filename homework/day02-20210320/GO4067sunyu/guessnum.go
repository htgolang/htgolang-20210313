package main

import "fmt"

func main() {
	slice := []int{9, 8, 19, 10, 2, 8}
	for i := 0; i < len(slice); i++ {
		for k := 0; k < len(slice)-1; k++ {
			if slice[k] > slice[k+1] {
				slice[k], slice[k+1] = slice[k+1], slice[k]
			}
		}
	}
	fmt.Println(slice)
	max := len(slice)
	mid := (max - 1) / 2
	tag := false
	var input int
	fmt.Print("请输入[9, 8, 19, 10, 2, 8]中的数字：")
	fmt.Scan(&input)
	if slice[mid] == input {
		fmt.Printf("输入数字的索引是%v", mid)
	} else if slice[mid] > input {
		for k := 0; k < mid; k++ {
			if slice[k] == input {
				fmt.Printf("输入数字的索引是%v", k)
				tag = true
			}
		}
	} else if slice[mid] < input {
		for k := mid + 1; k < len(slice); k++ {
			if slice[k] == input {
				fmt.Printf("输入数字的索引是%v", k)
				tag = true
			}
		}
	}
	if tag == false {
		fmt.Print("不存在")
	}
}
