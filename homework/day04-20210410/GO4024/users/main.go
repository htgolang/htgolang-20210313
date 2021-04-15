package main


import (
	"strings"
	//"os/user"
	"fmt"
	"github.com/princebot/getpass"
	"crypto/md5"
	_ "users/user"
	cc "users/commands"
)

func input(prompt string) string {
	fmt.Println(prompt)
	var txt string
	fmt.Scan(&txt)
	return txt
}

func prompt() {
	prompts := cc.Prompts
	fmt.Println(strings.Repeat("*",10))
	fmt.Println("欢迎使用xxx管理系统")
	fmt.Println("0. 退出")
	for _,v := range prompts {
		fmt.Printf("%s. %s\n",v[0],v[1])
	}
	//fmt.Println(strings.Repeat("*",10)
}

const Password = "123456"
func main() {
	//启动侯给欢迎
	//输入密码 3次输入错误 退出
	
	//输入密码
	// var password string 
	// fmt.Println("请输入密码:")
	// fmt.Scan(&password)
	
	//getpass 第三方函数
	//fmt.Println(getpass.Get("请输入密码:"))
	
	for i := 0; i < 3; i ++ {
	password, _ := getpass.Get("请输入密码:")
	//fmt.Println("你输入的密码:",string(password))
	//fmt.Printf("%T",password)
	if md5.Sum(password) != md5.Sum([]byte(Password)) {
		fmt.Println("密码错误，请重新输入:")
		
	} else {
		fmt.Println("密码正确")
		break
	}
}

//提示用户输入指令： 1，添加，2，编辑，3，删除，4，查询，5，退出

	commands := cc.Commands
END:
for {
	prompt()
	cmd := input("请输入指令")
	if command, ok := commands[cmd]; ok{
		command()
	} else if cmd == "0" {
		break END
	} else {
		fmt.Println("指令错误")
	}
	
}

}