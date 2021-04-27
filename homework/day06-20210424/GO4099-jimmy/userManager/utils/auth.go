package utils

import (
	"fmt"
	"github.com/princebot/getpass"
	"log"
	"os"
	"time"
	"userManager/model"
)

const (
	MaxAuthCount = 3
	LogPath = "./log/user.log"
)

var CurrentLoginUserId int

func Auth(users map[int]model.User) bool {
	flag := false
	username := Input("请输入用户名: ")
	for _, user := range users {
		if user.Name == username {
			// 密码验证
			for i := 0; i < MaxAuthCount; i++ {
				if user.Name == username {
					pwd, err := getpass.Get("请输入密码: ")
					if err != nil {
						log.Fatal(err)
					}
					if string(pwd) == user.Password {
						flag = true
						CurrentLoginUserId = user.Id
						file ,err := os.OpenFile(LogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatal(err)
						}
						defer file.Close()
						file.WriteString(time.Now().Format("2006-01-02 15:04:05") + " " + user.Name+ "用户登录成功\n")
						return flag
					} else {
						if i == MaxAuthCount-1 {
							fmt.Println("密码错误，您的三次机会已用完")
							return false
						} else {
							fmt.Printf("密码错误，您还有%d次机会\n", MaxAuthCount-i-1)
						}
					}
				}
			}
		}
	}
	if flag != true {
		fmt.Println("用户不存在")
	}
	return flag
}
