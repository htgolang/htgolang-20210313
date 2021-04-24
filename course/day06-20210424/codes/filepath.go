package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Abs("."))
	path := "/root/kk.txt"
	fmt.Println(filepath.Dir(path), filepath.Base(path), filepath.Ext(path))

	fmt.Println(filepath.IsAbs("."), filepath.IsAbs("/root/kk.txt"), filepath.IsAbs("c:/"))

	fmt.Println(filepath.Clean("/root/a/../kk.txt"))
	fmt.Println(filepath.Clean("/root/a/../kk.txt"))
	fmt.Println(filepath.FromSlash("/root/a/..\\kk.txt"), filepath.ToSlash("/root/a/..\\kk.txt"))

	fmt.Println(filepath.Join("/root/kk", "a", "b", "kk.txt"))
	fmt.Println(filepath.Split(path))

	// /root/所有的.go文件 /root/*.go
	// glob
	// * 匹配任意多个
	// ? 匹配一个
	// (f1|f2|f3)

	fmt.Println(filepath.Glob("./*.go"))
	fmt.Println(filepath.Glob("./*.txt"))
	fmt.Println(filepath.Glob("./*/*.txt"))
	fmt.Println(filepath.Match("./b/*.txt", "./b/a.txt"))
	fmt.Println(filepath.Match("./b/*.txt", "./b/atxt"))

	filepath.Walk(".", func(path string, fileInfo fs.FileInfo, err error) error {
		fmt.Println(path, fileInfo.IsDir(), fileInfo.Name())
		return nil
	})

}
