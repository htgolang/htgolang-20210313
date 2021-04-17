package common

import (
	"crypto/md5"
	"fmt"
	"github.com/princebot/getpass"
	"strconv"
	"strings"
	"time"
	"userManager/model"
)

const (
	MaxAuthCount = 3
	Password     = "e10adc3949ba59abbe56e057f20f883e"
)

func Input(prompt string) string {
	var text string
	fmt.Print(prompt)
	fmt.Scan(&text)
	return strings.TrimSpace(text)
}

// 密码验证
func Auth() bool {
	for i := 0; i < MaxAuthCount; i++ {
		pwd, err := getpass.Get("请输入密码: ")
		if err != nil {
			panic(err)
		}
		if fmt.Sprintf("%x", md5.Sum(pwd)) == Password {
			return true
		} else {
			if i == MaxAuthCount-1 {
				fmt.Println("密码错误，您的三次机会已用完")
			} else {
				fmt.Printf("密码错误，您还有%d次机会\n", MaxAuthCount-i-1)
			}
		}
	}
	return false
}

func InputUser(id int) *model.User {
	name := Input("请输入用户名: ")
	tel := Input("请输入电话: ")
	addr := Input("请输入地址: ")
	return model.NewUser(id, name, addr, tel)
}

func InputTodo(id int) *model.Todo {
	name := Input("请输入名称: ")
	priority := Input("请输入优先级: ")
	desc := Input("请输入说明: ")
	schedule := Input("请输入进度: ")
	status, _  := strconv.Atoi(Input("请输入状态(0/1): "))
	principal := Input("请输入负责人: ")
	startAt := Input("请输入开始时间: ")
	endAt := Input("请输入结束时间: ")
	completeAt := Input("请输入完成时间: ")
	return model.NewTodo(id, name, priority, desc, schedule, status, principal, strToTime(startAt), strToTime(endAt), strToTime(completeAt))
}

func strToTime(t string) time.Time {
	tm, err := time.Parse("2006-01-02", t)
	if err !=nil {
		fmt.Println(err)
	}
	return tm
}
