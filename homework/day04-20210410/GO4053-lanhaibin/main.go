package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"user/command"
	_ "user/init"

	"github.com/howeyc/gopass"
)

const password = "e10adc3949ba59abbe56e057f20f883e"

func askpass() {
	fmt.Print("请输入密码:")
	bytes, _ := gopass.GetPasswd()
	md5sum := md5.New()
	io.WriteString(md5sum, string(bytes))
	if fmt.Sprintf("%x", md5sum.Sum(nil)) != password {
		fmt.Println("密码错误!")
		os.Exit(1)
	}
}

func prompt() {
	fmt.Println("请输入要执行的操作:")
	for _, v := range command.Prompts {
		fmt.Printf("%s: %s\n", v[0], v[1])
	}
}

func main() {
	askpass()
	var op string
	for {
		prompt()
		fmt.Print("请输入要执行的操作:")
		fmt.Scan(&op)
		f, ok := command.Commands[op]
		if ok {
			f()
		}
	}
}
