package main

import "fmt"

// 统计一个字符串中每个单词出现的次数。例如："Hello, how do you do?"中，how=1,do=2,you=1

func main() {

	//定义silce
	var silce1 = []int{2,8,9,10,8,19}

	//假设最大的元素是silce[0]
	largest :=silce1[0]

	for i:=1;i<len(silce1);i++{
		//遍历silce，如果silce[i]大于largest，则最大值是为silce[i]
		if largest<silce1[i]{

			largest = silce1[i]

		}

	}

	//根据最大值找索引
	for j:=1;j<len(silce1);j++{
		//遍历silce，如果silce[i]大于largest，则最大值是为silce[i]
		if largest==silce1[j]{

			fmt.Println(j,largest)

		}

	}

}
