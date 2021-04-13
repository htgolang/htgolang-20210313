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
func Md5sum(txt string) (md5Value string) {
	md5Value = fmt.Sprintf("%X", md5.Sum([]byte(txt)))
	return
}

// 帮助信息
func HelpPrompt() {
	str := `--------------------------------------------------
	***********用户管理系统***********
	    1    ）添加用户
  	    2    ）修改用户
	    3    ）删除用户
	    4    ）搜索信息
	    5    ）打印用户列表
	    help ) 帮助信息
	    0    ）退出系统
	********************************
	`
	fmt.Println(str)
}
