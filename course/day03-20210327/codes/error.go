package main

import (
	"errors"
	"fmt"
)

/*
错误：
	语法错误 => 编译失败
	运行时 =>
		可恢复的错误/不可恢复的错误

		网络问题失败
		读取文件失败
		写文件磁盘满了
		...
		...
		...
*/

func fact(n int64) (int64, error) {
	if n < 0 {
		return 0, errors.New("计算错误")
	}
	if n == 1 || n == 0 {
		return 1, nil
	}
	r, err := fact(n - 1)
	if err == nil {
		return n * r, nil
	}
	return 0, fmt.Errorf("计算错误")
}

func main() {
	fmt.Println(fact(-3))
}
