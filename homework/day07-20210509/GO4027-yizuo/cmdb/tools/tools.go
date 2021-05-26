package tools

import (
	"crypto/md5"
	"fmt"
)

// 确认是否执行
func UserInputExecuted(data string) bool {
	userInputExecutedData := StrInput(data)
	switch userInputExecutedData {
	case "y", "Y", "yes", "YES":
		return true
	default:
		return false
	}
}

// 传入字符串，自动返回用户输入的值
func StrInput(strData string) (strInput string) {
	fmt.Println(strData)
	fmt.Scan(&strInput)
	return strInput
}

// 将传递的string转换为MD5加密后的十六进制返回
func Md5sum(str string) (md5Value string) {
	md5Value = fmt.Sprintf("%X", md5.Sum([]byte(str)))
	return
}

// 字符串检测验证，超过次数退出
func StringInputCheck(str, passwd string, num int) bool {
	/*
		循环三次，对比用户输入的用户名密码是否正确
			密码正确返回 true
			密码错误返回 false
	*/
	for i := 0; i < num; i++ {
		if i != 0 {
			fmt.Println(`验证错误，请重新输入：`)
		}
		inputPasswd := StrInput(str)

		// 对比输入值是否与预设值一致
		if inputPasswd == passwd {
			return true
		}
	}
	return false
}
