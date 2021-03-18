package main

import "fmt"

// 杨辉三角

func main() {
	var n int
	fmt.Printf("请输入层数:")
	fmt.Scanf("%d", &n)
	for i:=1;i<=n;i++ {
		for j,k:=1,1;j<=i;j++ {
			fmt.Printf("%d ", k)
			k = k * (i-j) / j
		}
		fmt.Println()
	}
}