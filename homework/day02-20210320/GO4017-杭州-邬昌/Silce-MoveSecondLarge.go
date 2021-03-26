package main

import "fmt"

// 统计一个字符串中每个单词出现的次数。例如："Hello, how do you do?"中，how=1,do=2,you=1

func main() {

	//定义silce
	var silce1 = []int{2,8,9,19,10,8,7,12,3,14}

	for i:=0;i<len(silce1)-1;i++{
		//遍历silce，如果左边大于右边，则两数调换位置
		if silce1[i] >=silce1[i+1]{

			silce1[i],silce1[i+1]=silce1[i+1],silce1[i]

		}

		//fmt.Println(silce1)


	}

	//再排序一遍，把第二大的排到倒数第二位置
	for j:=0;j<len(silce1)-1;j++{

		if silce1[j] >silce1[j+1]{

			silce1[j],silce1[j+1]=silce1[j+1],silce1[j]

		}
	}
	fmt.Println(silce1)

}
