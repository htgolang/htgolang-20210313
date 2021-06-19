package utils

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

// 传入字符串，自动返回用户输入的值
func StrInput(strData string) (strInput string) {
	fmt.Println(strData)
	fmt.Scan(&strInput)
	return strInput
}

// 将string转换为int
func StrconvAtoi(s string) (i int) {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("String 转换为 int 失败", err)
	}
	return i
}

// 将int转换为string
func StrconvItoa(i int) (s string) {
	return strconv.Itoa(i)
}

// 将传递的string转换为MD5加密后的十六进制返回
func Md5sum(str string) (md5Value string) {
	md5Value = fmt.Sprintf("%X", md5.Sum([]byte(str)))
	return
}

// 如果输入的值是"男"or"ture"，则返回一个bool值
func TextSex(t string) bool {
	if t == "男" || t == "true" {
		return true
	}
	return false
}
