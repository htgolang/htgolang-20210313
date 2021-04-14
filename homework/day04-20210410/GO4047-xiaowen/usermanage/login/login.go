package login

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"github.com/princebot/getpass"
)



func Login(){
	var num = 1
	const pwd = `202cb962ac59075b964b07152d234b70`
	pw, err := getpass.Get("请输入密码: ")
	for {
		//得到密码MD5的字符串
		tmp := md5.New()
		tmp.Write(pw)
		tmppw := hex.EncodeToString(tmp.Sum(nil))
		if err != nil {
			fmt.Println(err)
		} 
		//验证
		if tmppw == pwd {
			break
		} else{
			fmt.Println("你输入的密码有误！！")
			num++
			pw, _ = getpass.Get("请重新输入密码: ")
		}
		if num == 3{
			os.Exit(1)
		}
	}
}