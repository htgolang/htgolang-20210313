package utils

import (
	"fmt"
	"github.com/princebot/getpass"
	"sort"
	"userManager/config"
)

func VerifyPWD() (err error) {
	defer func() {
		if errMsg := recover(); errMsg != nil {
			err = fmt.Errorf("%v", errMsg)
		}
	}()

	times := 3
	for times != 0 {
		pw, ok := getpass.Get("输入密码:")

		if GetMD5(string(pw)) == config.PAS && ok == nil {
			fmt.Println("登录成功")
			return
		}
		times--
		fmt.Printf("密码错误, 还有%d次机会\n", times)
	}
	panic("用户信息验证失败, 退出程序")
}

func ShowMenu(choice string) {
	if choice == "1" {
		fmt.Println("\n\n用户TODO管理菜单")
		sortSlice := make([]string, 0, len(config.TodoMenuMap))
		for k := range config.TodoMenuMap {
			sortSlice = append(sortSlice, k)
		}
		sort.Strings(sortSlice)
		for _, k := range sortSlice {
			fmt.Printf("%s          %s\n", k, config.TodoMenuMap[k])
		}
	} else {
		fmt.Println("\n\n用户信息管理菜单")
		sortSlice := make([]string, 0, len(config.UserMenuMap))
		for k := range config.UserMenuMap {
			sortSlice = append(sortSlice, k)
		}
		sort.Strings(sortSlice)
		for _, k := range sortSlice {
			fmt.Printf("%s          %s\n", k, config.UserMenuMap[k])
		}
	}
}
