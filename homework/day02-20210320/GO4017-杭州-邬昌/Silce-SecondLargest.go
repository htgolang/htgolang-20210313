package main

import "fmt"

// 统计一个字符串中每个单词出现的次数。例如："Hello, how do you do?"中，how=1,do=2,you=1

func main() {

	//定义silce
	var silce1 = []int{2,8,9,19,10,8,7,12,3,14}

	//假设最大的元素是silce[0]，第二大是silce[1]
	largest :=silce1[0]
	secondLargest :=silce1[1]

	for i:=1;i<len(silce1);i++{
		//遍历silce，如果silce[i]大于largest，则最大值是为silce[i]
		if largest<silce1[i]{
			//将上一个最大值赋值给第二大的值
			secondLargest =largest
			largest = silce1[i]
        //判断第二大的值小于最大值前提下，
		}else if secondLargest <largest && secondLargest < silce1[i]{

			secondLargest =silce1[i]
		}

	}

	fmt.Println(secondLargest)

}
