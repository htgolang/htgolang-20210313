package main

import (
	f "fmt"       // 别名导入
	. "testmod/a" // 点导入 调用包内函数变量 不加包名
	"testmod/b"

	"testmod/a/math"   // 绝对路径导入 => 相对路径导入 gomod 禁用
	_ "testmod/b/math" // 下划线导入 // 包初始化
)

func main() {
	// a.Name, a.Call
	f.Println(Name)

	Call()

	// b.Name b.Call
	f.Println(b.Name)
	b.Call()

	// a.math
	f.Println(math.PI)
	f.Println(math.Add(1, 2))

	// b.math
	//f.Println(bmath.PI, bmath.Add(1, 2))

	f.Println(Version)

}
