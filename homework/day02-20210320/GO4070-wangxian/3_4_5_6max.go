//找到切片中最大的元素索引和值(打印)

package main

import "fmt"

func main() {
	var s1 = []int{111, 22, 34, 99, 8, 4, 10}
	//找到切片中最大的元素索引和值(打印)
	var max_num int
	var max_index int

	for i, v := range s1 {
		if v > max_num {
			max_num = v
			max_index = i
		}
	}
	fmt.Printf("最大值是%d,索引是%d\n", max_num, max_index)

	//找到切片第二大的元素索引和值(打印)
	var second_num int
	var second_index int
	for i, v := range s1 {
		if i == max_index {
			continue
		} else if v > second_num {
			second_num = v
			second_index = i
		}
	}
	fmt.Printf("第二大是%d,索引是%d\n", second_num, second_index)

	//将最大的元素移动到切片最末尾
	s1[max_index], s1[len(s1)-1] = s1[len(s1)-1], s1[max_index]
	fmt.Println(s1)

	//将第二大的元素移动到切片倒数第二位
	s1[second_index], s1[len(s1)-2] = s1[len(s1)-2], s1[second_index]
	fmt.Println(s1)

}
