package view
import (
	"github.com/princebot/getpass"
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	_ "mgr/model/init"
	"mgr/controller/user"
	"mgr/controller/commands"
	"mgr/controller/utils"
)


func prompt() {
	fmt.Println(strings.Repeat("*", 31))
	fmt.Println(" ====欢迎使用多云管理平台====")
	fmt.Println("    1. 退出")
	commands.Commands.Prompt(2, "    %d. %s\n")
	fmt.Println(strings.Repeat("*", 31))
}

func userLogin() bool{

	fmt.Printf("请输入用户名:")
	fmt.Scan(&user.CurrentUser)
	passwordBytes, _ := getpass.Get("请输入密码:")
	md5Pwd := fmt.Sprintf("%x", md5.Sum([]byte(passwordBytes)))

	str := ""
	var users = user.Users
	for _, u := range users {
		if user.CurrentUser == u.Name{
			str = u.Pwd
		}
	}
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	if md5str == md5Pwd {
		fmt.Printf("恭喜%v登录成功!\n", user.CurrentUser)
		return true
	}else {
		return false
	}
}


func Login(){
	count := 0
	for count < 3 {
		flag := userLogin()
		if flag {
			break
		} else {
			count = count + 1
			if count == 3 {
				fmt.Printf("输入次数已经达%v,将退出程序!", count)
				return
			} else {
				fmt.Println("用户名或密码不正确,请重新输入!")
				continue
			}
		}
	}

END:
	for {
		prompt()
		cmd, _ := strconv.Atoi(utils.Input("请输入指令:"))
		if command := commands.Commands.Get(cmd); command != nil {
			command(user.Users)
		} else if cmd == 1 {
			break END
		} else {
			fmt.Println("指令错误")
		}
	}
}
