//打印杨辉三角形

package main

import (
	"fmt"
)

func main() {
	YhTriangle()

}

const lines int = 10

func YhTriangle() {
	var YhT [lines][lines]int
	for i := 0; i < lines; i++ {

		for j := 0; j < i+1; j++ {
			if i < 2 { //两行以内三角形的数字都是1
				YhT[i][j] = 1

			} else {
				//第三行开始，正式计算数值写入数组
				if j == 0 || j == i {
					YhT[i][j] = 1 //所有行的第一列和最后一列都是1
				} else {
					YhT[i][j] = YhT[i-1][j-1] + YhT[i-1][j] //当前数组元素是上一行的前一个元素加上上一行的当前列元素
				}
			}
			fmt.Printf("%d\t", YhT[i][j]) //格式化输出一行
		}
		fmt.Print("\n")
	}
}
