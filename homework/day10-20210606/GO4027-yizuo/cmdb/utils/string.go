package utils

import (
	"crypto/md5"
	"encoding/json"
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

// 如果输入的值是bool值,则返回一个 "男"or"女"，
func SexText(b bool) string {
	if b {
		return "男"
	}
	return "女"
}

// 结构体转换为json格式的字符串并返回
func StructJsonMarshal(i interface{}) string {
	data, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// 判断值是否全部是字符串或者数字(没有特殊字符)
func IsLetterOrNumer(txt string) bool {
	for _, v := range txt {
		if !(v >= 'a' || v >= 'A' || v >= '0' ||
			v <= 'z' || v <= 'Z' || v <= '9') {
			return false
		}
	}
	return true
}
