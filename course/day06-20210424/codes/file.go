package main

import (
	"fmt"
	"os"
)

func main() {
	// 文件
	// 1. 删除
	fmt.Println(os.Remove("test.a"))
	// 2. 文件重命名
	fmt.Println(os.Rename("test.b", "test.c"))
	// 3. 设置权限
	os.Chmod("test.c", 0600)
	// 目录
	os.Mkdir("xxx", os.ModePerm)
	os.MkdirAll("x/y", os.ModePerm)
	os.Remove("xxx")
	fmt.Println(os.Remove("x"))
	os.RemoveAll("x")
	os.Rename("a", "b")

	// fmt.Println(os.Getwd())
	// os.Chdir("d://")
	// fmt.Println(os.Getwd())

	// os.Create("xxx.txt")

}
