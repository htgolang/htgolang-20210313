package utils

import "strings"

//判断cookie中sid值是否合法
func IsLetterAndNum(txt string) bool {
	s := strings.ReplaceAll(txt, "-", "")
	for _, ch := range s {
		if !(ch >= 'a' && ch <= 'z' ||
			ch >= 'A' && ch <= 'Z' ||
			ch >= '0' && ch <= '9') {
			return false
		}
	}
	return true
}
