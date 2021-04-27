package handles

import (
	"fmt"

	"github.com/princebot/getpass"
)

func CheckPassword() bool{

	var flag int
	const pass = "123456"

	for {

		//fmt.Scan(&password)
		password, _ := getpass.Get("请输入密码：")
		if string(password) == pass {

			fmt.Println("欢迎，请继续操作！！！")
			break

		} else {

			flag++
			fmt.Println("密码错误，请重新输入密码！！！")

			if flag > 2 {

				fmt.Println("密码错误3次了，不能再尝试！！！！")

				return false

			}
		}
	}

	return true
}
