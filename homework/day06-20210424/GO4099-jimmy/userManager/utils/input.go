package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"userManager/model"
)

func Input(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func InputUser(id int, operating, pwd string) model.User {
	name := Input("请输入用户名: ")
	tel := Input("请输入电话: ")
	addr := Input("请输入地址: ")
	if operating == "add" {
		pwd = Input("请输入密码: ")
	}

	return model.NewUser(id, name, addr, tel, pwd)
}
