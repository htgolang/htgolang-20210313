package main

import "fmt"

func main() {
	yanghuiTriangle(5)
}
/*杨辉三角形
百度百科的描述：https://baike.baidu.com/item/%E6%9D%A8%E8%BE%89%E4%B8%89%E8%A7%92/215098?fr=aladdin
*/
func yanghuiTriangle (num int) {
	//如果没指定层数，默认打印五层
	if num == 0 {
		num = 5
	}
	a := make([][]int,0) //创建二维数组，长度为0
	a = append(a,[]int{1}) //切片左侧都是1
	for i := 1; i < num; i++ {
		temp := []int{1}  
		for j := 1; j <= i-1; j++{
			temp = append(temp,a[i-1][j-1]+a[i-1][j]) //当前数值是上一行的前一个值加上上一行的当前列元素		
		}
		temp = append(temp,1) //切片右侧都是1
		//fmt.Println(temp)
		a = append(a,temp) //加上初始的第一行 
	}
	//打印杨辉三角形
	for k := 0; k < len(a); k++ {
		for m := 0; m < len(a[k]); m++ {
		    fmt.Printf("%d\t",a[k][m])
		}
		fmt.Println()
	}	
}