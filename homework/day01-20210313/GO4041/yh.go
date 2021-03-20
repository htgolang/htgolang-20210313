package main

import "fmt"

func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)

		// 首尾为1
		ans[i][0] = 1
		ans[i][i] = 1

		// 遍历每一行数组内的每一个元素=上一行左右2个数字相加
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}

	}
	return ans
}

func main() {
	fmt.Println("请输入打印的阶数：")

	var k int
	fmt.Scanf("%d", &k)
	fmt.Println("杨辉三角：")
	yh := generate(k)
 
	for i, j := range yh {
		for k := 0; k <= len(yh)-i; k++ {
			fmt.Printf("  ")
		}
		for _, l := range j {
			fmt.Printf("%3d ", l)
		}

		fmt.Println()
	}

}
