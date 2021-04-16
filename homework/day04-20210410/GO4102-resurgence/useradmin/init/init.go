package init

import (
	"fmt"
	"github.com/princebot/getpass"
	"os"
	. "useradmin/config"
	. "useradmin/user"
	"useradmin/utils"
)

func init() {
	if err := verifyPWD(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	Register("显示用户信息", GetUser)
	Register("添加用户信息", AddUser)
	Register("更新用户信息", UpdateUser)
	Register("删除用户信息", DeleteUser)
	Register("退出", func(i *[]*User) error {
		os.Exit(0)
		return nil
	})
}

func verifyPWD() (err error) {
	defer func() {
		if errMsg := recover(); errMsg != nil {
			err = fmt.Errorf("%v", errMsg)
		}
	}()
	times := 3
	for times != 0 {
		pw, _ := getpass.Get("输入密码:")

		if utils.GetMD5(string(pw)) == PAS {
			fmt.Println("登录成功")
			return
		}
		times--
		fmt.Printf("密码错误, 还有%d次机会\n", times)
	}
	panic("用户信息验证失败, 退出程序")
}
