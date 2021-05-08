package handles

import (
	"crypto/md5"
	"fmt"
	"github.com/princebot/getpass"
	"io/ioutil"
	"os"
)

func CheckPassword() bool{

	var flag int
	const passwd = "passwordfile"

	for {

		//fmt.Scan(&password)
		pass, err := ioutil.ReadFile(passwd)
		if err ==nil{

			bytes, _ := getpass.Get("请输入密码：")
			if fmt.Sprintf("%X",md5.Sum([]byte(bytes))) == string(pass) {

				fmt.Println("欢迎，请继续操作！！！")
				return true

		} else {

				flag++
				fmt.Println("密码错误，请重新输入密码！！！")

				if flag > 2 {

					fmt.Println("密码错误3次了，不能再尝试！！！！")

					return false

				}
			}

		}else {
			if os.IsNotExist(err){
				bytes,_ :=getpass.Get("请输入初始化密码：")
				ioutil.WriteFile(passwd,[]byte(fmt.Sprintf("%X",md5.Sum(bytes))),os.ModePerm)
				return true

			}else {

				fmt.Println(err)
				return false
			}
		}
	}

	return true
}
