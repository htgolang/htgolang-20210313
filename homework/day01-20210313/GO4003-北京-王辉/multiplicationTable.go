package main

import "fmt"
/*
打印乘法口诀
1 1 =>1
2 * 1 => 2 2 * 2=>4
9 * 1 => 9 .....    9*9 => 81
 */
func main() {
	for row:=1;row<10;row++ {
		for col:=1;col<=row;col++ {
			fmt.Printf("%d*%d=%d\t",col,row,col*row)
		}
		fmt.Println()
	}
}
