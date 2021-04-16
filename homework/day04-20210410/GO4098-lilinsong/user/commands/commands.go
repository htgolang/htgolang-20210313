package commands

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/princebot/getpass"
	"os"
	"strconv"
	co "user/consoles"
)

type info struct {
	addr, tel string
}

var account = make(map[string]*info, 0)

var auth = "e10adc3949ba59abbe56e057f20f883e"

func addUser(user, addr, tel string) map[string]*info {
	if ok := doesUserExist(user); ok {
		return nil
	}
	msg := info{
		addr: addr,
		tel:  tel,
	}
	account[user] = &msg

	return account
}

func doesUserExist(user string) bool {
	if _, ok := account[user]; ok {
		fmt.Println("用户以存在")
		return true
	} else {
		fmt.Println("用户不存在")
		return false
	}
}

func Add() {
	user := Input("请输入用户的姓名: ")
	if ok := doesUserExist(user); ok {
		return
	}

	addr := Input("请输入用户的地址: ")
	phone := Input("请输入用户的电话: ")

	addUser(user, addr, phone)
}

func Del() {
	user := Input("请输入用户的姓名: ")
	if ok := doesUserExist(user); ok {
		delete(account, user)
		fmt.Println("删除用户成功")
	}
}

func Query() {
	text := Input("请输入要查询用户的姓名: ")
	if ok := doesUserExist(text); ok {
		fmt.Println(text, account[text].addr, account[text].tel)
	}
}

func Modify() {
	user := Input("请输入要修改用户的名称: ")
	if ok := doesUserExist(user); ok {
		text := Input("请输入要修改的字段(user/addr/tel): ")
		switch text {
		case "user":
			content := Input("请输入用户昵称: ")
			account[content] = account[user]
		case "addr":
			content := Input("请输入新的地址: ")
			account[user].addr = content
		case "tel":
			content := Input("请输入新的电话: ")
			account[user].tel = content
		default:
			fmt.Println("输入的字段不存在")
			return
		}
	}

}

func md5Hash(pass string) error {

	hash := md5.New()
	hash.Write([]byte(pass))
	password := fmt.Sprintf("%x", string(hash.Sum(nil)))

	if password == auth {
		return nil
	} else {
		return errors.New("密码输入错误")
	}
}

func Login() bool {
	var count = 0
	for {
		if count >= 3 {
			return false
		}
		pass, _ := getpass.Get("password: ")
		if err := md5Hash(string(pass)); err != nil {
			fmt.Printf("\nPermission denied, please try again.\n")
			count++
			continue
		} else {
			return true
		}
	}
}

func Entrance() {
	writer := tablewriter.NewWriter(os.Stdout)
	writer.SetHeader([]string{"选项", "说明"})
	for i := 1; i <= len(co.Menu); i++ {
		writer.Append([]string{
			strconv.Itoa(i), co.Menu[i],
		})
	}
	writer.Render()
}

func MenuSelection() {
	text := Input("请输入你要输入的序号: ")
	num, _ := strconv.Atoi(text)
	if num == 1 {
		fmt.Println("-_-!!!")
		os.Exit(0)
	}
	if value, ok := co.ManagementView.Get(num); !ok {
		fmt.Printf("输入有误,序号范围%d~%d\n", 1, co.ManagementView.Len()+1)
	} else {
		fn := value.(func())
		fn()
	}
}

func Input(content string) string {
	var text string
	fmt.Print(content)
	fmt.Scan(&text)
	return fmt.Sprintf("%s", text)
}
