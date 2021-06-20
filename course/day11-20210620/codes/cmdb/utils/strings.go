package utils

func IsLetterOrNumer(txt string) bool {
	for _, ch := range txt {
		if !(ch >= 'a' && ch <= 'z' ||
			ch >= 'A' && ch <= 'Z' ||
			ch >= '0' && ch <= '9') {
			return false
		}
	}
	return true
}
