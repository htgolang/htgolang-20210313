package utils

func IsLetterOrNumber(txt string) bool{
	for _, v := range txt{
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' || v >= '0' && v <= '9'{
			return true
		}
	}
	return false
}