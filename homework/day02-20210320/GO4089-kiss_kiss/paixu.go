package main

import (
	"fmt"
	"strconv"
)

/*
思路，比较相邻的两个值，比较大小，如果第一个比第二个大，就交换他们两个。
	外循环：第一次排序
	8,9,19,10,2,8
	8,9,19,10,2,8
	8,9,10,19,2,8
	8,9,10,2,19,8
	8,9,10,2,8,19
	最大的就会得出来，并放到最后面
	所以第二次循环最后一个就不需要再比较了
*/
func main() {
	var (
		aa       = []int{}
		loopback int
		input    string
	)
	for {
		fmt.Print("请输入数字(exit退出)：")
		fmt.Scan(&input)
		if input == "exit" {
			break
		} else {
			UserInput, er := strconv.Atoi(input)
			if er == nil {
				aa = append(aa, UserInput)
			} else {
				fmt.Println("输入的类型非整数，请重新输入")
				continue
			}
		}
	}
	//fmt.Println("--------", aa, "----------")
	//冒泡
	for i, _ := range aa {
		for j, _ := range aa {
			//正序 aa[:len(aa)-i-1]
			//倒叙 aa[len(aa)-i-1:]
			//for j, _ := range aa[len(aa)-i-1:] {
			if aa[i] < aa[j] { //正序aa[i] < aa[j]  倒叙 aa[i] > aa[j]
				aa[i], aa[j] = aa[j], aa[i]
				loopback++
			}
		}
	}
	fmt.Println(aa, loopback)
	//遍历查找
	// for {
	// 	fmt.Print("请输入要查找的数字(exit退出)：")
	// 	fmt.Scan(&input)
	// 	if input == "exit" {
	// 		break
	// 	} else {
	// 		//查找的主要代码块
	// 		var Tage int
	// 		if InputValue, er := strconv.Atoi(input); er == nil {
	// 			for Index, Value := range aa {
	// 				if Value == InputValue {
	// 					Tage = 1
	// 					fmt.Println(Index)
	// 				}
	// 			}
	// 			if Tage == 0 {
	// 				fmt.Println(-1)
	// 			}
	// 			break
	// 		} else {
	// 			fmt.Println("输入的类型非整数，请重新输入!")
	// 			continue
	// 		}
	// 	}
	// }
	/*二分查找
	思路
			a.第一次比较中间的数值是否与输入的数值相等，如果相等，返回索引值（索引值为tag+len(aa)/2，tag初始化为0）
			b.如果不等，则比较边缘左右值：
				若左边缘值大于右边缘值，说明序列时从大到小排列
				6，5，4，3，2，1
				比较中间值和输入值的大小
					b1.若输入值小于中间值（假设输入值为1）
						输入值1小于中间值4，则说明值不在左半部分，将左半部分去除，此时新切片（aa=aa[len(aa)/2+1:]）为3，2，1
						此时3的索引值由3变为0，新旧索引的关系tag为：tag=len(aa)/2+1+tag
						重新进入a步骤进行比较
					b2.若输入值大于中间值（假设输入值6）
						6大于中间值4，则说明值不在右半部分，将右半部分去除，此时新切片（aa=aa[:len(aa)/2]）为6，5
						此时6和5的索引值未改变，所以tag的值不需要改变

	*/
	fmt.Print("请输入要查找的数字：")
	fmt.Scan(&input)
	if InputValue, er := strconv.Atoi(input); er == nil {
		Tag, label := 0, 0
		//tag
		for _ = range aa {
			if len(aa) < 1 {
				fmt.Println(-1)
				break
			} else if aa[len(aa)/2] == InputValue {
				fmt.Println("索引：", len(aa)/2+Tag)
				label = 1
				break
			} else if aa[0] < aa[len(aa)-1] {
				// 1，2，3，4，5
				if aa[len(aa)/2] > InputValue {
					aa = aa[:len(aa)/2]
				} else {
					Tag = len(aa)/2 + Tag
					aa = aa[len(aa)/2:]
				}
			} else {
				//6，5，4，3，2，1
				if aa[len(aa)/2] > InputValue {
					Tag = len(aa)/2 + Tag + 1
					aa = aa[len(aa)/2+1:]
				} else {
					aa = aa[len(aa)/2:]
				}
			}
		}
		if label == 0 {
			fmt.Println(-1)
		}
	} else {
		fmt.Println("输入的类型非整数，退出!")
	}
}
