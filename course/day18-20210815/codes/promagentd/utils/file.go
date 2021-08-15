package utils

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) string {
	if bytes, err := ioutil.ReadFile(path); err == nil {
		return string(bytes)
	}
	return ""
}

func WirteFile(path, content string) {
	ioutil.WriteFile(path, []byte(content), os.ModePerm)
}
