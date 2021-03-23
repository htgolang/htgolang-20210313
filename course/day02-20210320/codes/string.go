package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	txt := "我爱中华人民共和国"

	// fmt.Println(unsafe.Sizeof(txt))
	// 操作
	// 算数运算 +
	fmt.Println(txt + "!")
	// +=
	fmt.Println(txt)
	txt += "!"
	fmt.Println(txt)
	// 关系运算 > >= < <= != ==
	fmt.Println("abc" > "ad") // false
	// 获取字符串长度 len => 字节数量
	fmt.Println(len(txt)) // utf-8 1 2 3 4
	fmt.Println(utf8.RuneCountInString(txt))
	// 索引，切片
	fmt.Printf("%c\n", txt[27])
	// txt[27] = '%'
	fmt.Println(txt)
	// 遍历
	// for
	// for range
	for index := range txt {
		fmt.Println(index)
	}
	for i, v := range txt {
		fmt.Printf("%d: %T, %c\n", i, v, v)
	}
	// 切片 byte []byte
	// 切片 rune []rune
	bs := []byte(txt)
	rs := []rune(txt) // 长度字符数量
	fmt.Println(bs)
	fmt.Println(rs)

	fmt.Println(string(bs))
	fmt.Println(string(rs))

}
