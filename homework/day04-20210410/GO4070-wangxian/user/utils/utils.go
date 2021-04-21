package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

//md5加密
func Md5Encrypt(b []byte) string {
	hasher := md5.New()
	hasher.Write(b)
	hasher.Write([]byte("asxeft"))
	hash := fmt.Sprintf("%x", hasher.Sum(nil))
	return hash
}

//获取用户输入信息
func Input(msg string) (value string) {
	fmt.Print(msg)
	fmt.Scanln(&value)
	return strings.TrimSpace(value)
}

//检查用户是否存在,存在返回在切片中的索引，不存在返回-1
func CheckUserExist(id string, users []map[string]string) int {
	for i, v := range users {
		if id == v["id"] {
			return i
		}
	}
	return -1
}
