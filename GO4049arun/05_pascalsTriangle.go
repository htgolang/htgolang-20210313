package main

import "fmt"

func J(i int) int {
	var k int = 1
	for j := 1; j <= i; j++ {
		k = k * j
	}
	return (k)
}
func C(i int, j int) int { /*定义组合数*/
	var k int
	k = J(j) / (J(i) * J(j-i))
	return (k)
}
func main() {
	var i int /*打印杨辉三角*/
	fmt.Printf("请输入要打印的行数:")
	fmt.Scan(&i)
	if i <= 0 || i > 16 {
		fmt.Println("请输入[1,16]之间的数字!")
		return
	}
	fmt.Printf("%d行杨辉三角如下:\n", i)
	for j := 0; j < i; j++ {
		//输出前边的空格
		for k := 1; k <= (i - j); k++ {
			fmt.Printf(" ")
		}
		//输出杨辉三角数字
		for n := 0; n <= j; n++ {
			fmt.Printf("%v", C(n, j))
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}
