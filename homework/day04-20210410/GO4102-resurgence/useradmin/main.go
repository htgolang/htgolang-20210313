package main

import (
	"fmt"
	"strconv"
	. "useradmin/config"
	_ "useradmin/init"
	. "useradmin/utils"
)

func main() {

	var userSlice = make([]*User, 0, 100)

	for {
		showMenu()
		// 判断用户输入信息
		menu := Input("输入功能序号: ")
		if menu < "1" || menu > strconv.Itoa(len(FuncMap)) {
			fmt.Println("功能不合法, 重新输入")
			continue
		}
		if err := FuncMap[menu](&userSlice); err != nil {
			fmt.Println(err)
		}
	}
}

func showMenu() {
	fmt.Println("\n\n用户信息管理菜单")
	for _, arr := range MenuArr {
		fmt.Printf("%s          %s\n", arr[0], arr[1])
	}
}
