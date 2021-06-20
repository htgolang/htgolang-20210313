package main

import (
	_ "embed"
	"embedvars/pkg"
	"fmt"
)

//go:embed version
var version string

// 找文件目录 => 当前.go文件同级目录下

//go:embed version
var versionBytes []byte

func main() {
	fmt.Println("version: ", version)
	fmt.Println("versionBytes: ", versionBytes)
	fmt.Println("versionBytes: ", string(versionBytes))

	fmt.Println(pkg.Name)
	fmt.Println(pkg.Version)
}
