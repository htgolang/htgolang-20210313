package main

import "fmt"

/*
         1
        1 1
       1 2 1
      1 3 3 1
     1 4 6 4 1
    1 5 10 10 5 1
  1 6 15 20 15 6 1
 1 7 21 35 35 21 7 1
它的特征是，每一行的最左边和最右边的元素都是1，而其它元素等于它上方“肩膀”上的两个元素之和。
参考： https://www.runoob.com/note/27983
*/

func main() {
	var line int
	nums := []int{}
	fmt.Println("请输入需要打印多少行的杨辉三角形：")
	fmt.Scan(&line)
	for i :=0;i< line;i++ {
		for j:=0;j<(line-i);j++ {
			fmt.Print(" ")
		}
		for j :=0;j <(i+1);j ++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value =1
			}else {
				value = nums[length-i]+nums[length-i-1]
			}
			nums = append(nums,value)
			fmt.Print(value," ")
		}
		fmt.Println("")
	}
}