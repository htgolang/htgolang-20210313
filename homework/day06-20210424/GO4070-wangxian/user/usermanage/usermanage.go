package usermanage

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
	"user/module"
	"user/utils"
	"github.com/princebot/getpass"
	"gopkg.in/ini.v1"
)

var (
	CurrentUserIndex int                                                       //记录登录user
	Users            []module.User                                             //记录所有用户信息
	ConfPath         string        = "C:\\mygocode\\05\\user\\conf\\user.conf" //配置文件路径
	Confobj          *ini.File                                                 //加载后的配置文件
)

//加载ini配置文件
func LoadConfig(path string) *ini.File {
	ini.DefaultHeader = true
	conf, err := ini.Load(ConfPath)
	if err != nil {
		log.Fatalf("Load configure file failed")
	}
	return conf
}

//加载用户信息
func LoadUserInfo() []module.User {
	if Confobj.Section("DEFAULT").Key("persistent_type").String() == "gob" {
		path := filepath.Join(Confobj.Section("DEFAULT").Key("userdata_path").Value(), "user.gob")
		rs, err := GobPersisten{}.Load(path)
		if err != nil {
			log.Fatalf("load all user info faild. %s", err)
		}
		return rs
	} else {
		path := filepath.Join(Confobj.Section("DEFAULT").Key("userdata_path").Value(), "user.csv")
		rs, err := CsvPersistent{}.Load(path)
		if err != nil {
			log.Fatalf("load all user info faild. %s", err)
		}
		return rs
	}
}

//保存用户信息
func SaveUserInfo() error {
	var err error
	ptype := Confobj.Section("DEFAULT").Key("persistent_type").Value()
	path := filepath.Join(Confobj.Section("DEFAULT").Key("userdata_path").Value(), "user."+ptype)

	err = CleanOldFile()
	if err != nil{
		return err
	}

	if ptype == "gob" {
		err = GobPersisten{}.Save(Users, path)
	} else {
		err = CsvPersistent{}.Save(Users, path)
	}
	if err != nil {
		return err
	}
	return nil
}

//用户登录
func Login() bool {
	Confobj = LoadConfig(ConfPath)
	Users = LoadUserInfo()
	for i := 0; i < 3; i++ {
		name := utils.Input("请输入用户名:")
		pw, _ := getpass.Get("请输入密码:")

		for i, user := range Users {
			if name == user.Name && utils.Md5Encrypt(pw) == user.Passwd {
				//记录当前登录用户index
				CurrentUserIndex = i
				fmt.Printf("登录成功!\n\n")
				return true
			}
		}
		fmt.Printf("用户名或密码错误, 还可尝试%d次\n\n", 2-i)
	}
	return false
}

//添加user
func Add() {
	name := utils.Input("请输入名字:")
	addr := utils.Input("请输入地址:")
	tel := utils.Input("请输入电话:")
	pwd := utils.Input("请输入密码:")
	confirmPwd := utils.Input("请确认密码:")

	if CheckUserName(name) {
		fmt.Println("用户名已存在")
		return
	}

	if pwd != confirmPwd {
		fmt.Println("两次密码输入不一致")
		return
	}

	tmpuser := module.User{Id: generateID(), Name: name, Addr: addr, Tel: tel, Passwd: utils.Md5Encrypt([]byte(pwd))}
	Users = append(Users, tmpuser)

	fmt.Println(Users)

}

//删除user
func Delete() {
	id := utils.Input("请输入要删除的用户id:")
	index := CheckUserExist(id)
	if index != -1 {
		Users[index].Print()
		confirm := utils.Input("是否确认删除(y/n)?")
		switch confirm {
		case "y", "Y":
			Users = append(Users[:index], Users[index+1:]...)
			fmt.Println(Users)
		case "n":
			fmt.Println("取消删除")
		default:
			fmt.Println("输入错误")
		}
	} else {
		fmt.Println("用户不存在")
	}
}

//修改user
func Modify() {
	id := utils.Input("请输入要修改的用户id:")
	index := CheckUserExist(id)

	if index != -1 {
		Users[index].Print()
		confirm := utils.Input("是否确认修改(y/n)?")
		switch confirm {
		case "y", "Y":
			newName := utils.Input("请输入新的用户名:")
			newAddr := utils.Input("请输入新的地址:")
			newTel := utils.Input("请输入新的电话:")

			if CheckUserName(newName) {
				fmt.Println("用户名已存在")
				return
			} else {
				Users[index].Name = newName
				Users[index].Addr = newAddr
				Users[index].Tel = newTel
				fmt.Println(Users)
			}
		case "n", "N":
			fmt.Println("取消修改")
		default:
			fmt.Println("输入错误")
		}
	} else {
		fmt.Println("用户不存在")
	}

}

//查询user
func Query() {
	keyWord := utils.Input("请输入要查询的关键字:")

	flag := false
	for _, v := range Users {
		if strings.Contains(strconv.Itoa(v.Id), keyWord) || strings.Contains(v.Name, keyWord) || strings.Contains(v.Addr, keyWord) || strings.Contains(v.Tel, keyWord) {
			v.Print()
			flag = true
		}
	}

	if !flag {
		fmt.Println("没有匹配用户")
	}

}

//修改user密码
func ChangePasswd() {
	newpwd := utils.Input("请输入新的密码:")
	confirmpwd := utils.Input("请确认密码:")

	if newpwd == confirmpwd {
		Users[CurrentUserIndex].Passwd = utils.Md5Encrypt([]byte(newpwd))
		fmt.Println("密码修改成功")
		fmt.Println(Users[CurrentUserIndex])
	} else {
		fmt.Println("两次输入密码不一致，请稍后再试")
	}

}

//选择信息保存方式
func SelectPersistenType() {
	var Ptype string
	fmt.Print("请选择持久化方式(gob/csv): ")
	fmt.Scanln(&Ptype)
	switch Ptype {
	case "gob":
		Confobj.Section("DEFAULT").Key("persistent_type").SetValue("gob")
		Confobj.SaveTo(ConfPath)
	case "csv":
		Confobj.Section("DEFAULT").Key("persistent_type").SetValue("csv")
		Confobj.SaveTo(ConfPath)
	default:
		fmt.Println("选择错误,请选择gob或csv")
	}
}

//辅助功能函数
//生成用户ID
func generateID() int {
	id := 0
	for _, v := range Users {
		i := v.Id
		if i > id {
			id = i
		}
	}
	return id + 1
}

//检查用户名是否冲突
func CheckUserName(name string) bool {
	for _, u := range Users {
		if u.Name == name {
			return true
		}
	}
	return false
}

//检查用户是否存在,存在返回在切片中的索引，不存在返回-1
func CheckUserExist(uid string) int {
	id, _ := strconv.Atoi(uid)
	for i, v := range Users {
		if id == v.Id {
			return i
		}
	}
	return -1
}

//清理旧的用户信息文件，不论格式是gob还是csv，只保留最近的3个文件
func CleanOldFile() error {
	oldFilename := ""
	newFilename := ""
	path := Confobj.Section("DEFAULT").Key("userdata_path").Value()

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("clean old userinfo file faild, %v", err)
	}

	for _, f := range files {
		if f.Name() == "user.gob" {
			oldFilename = "user.gob"
		} else if f.Name() == "user.csv" {
			oldFilename = "user.csv"
		}
	}
	newFilename = oldFilename + "_" + time.Now().Format("2006-01-02-15-04-05")
	os.Rename(filepath.Join(path, oldFilename), filepath.Join(path, newFilename))

	files, err = ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("clean old userinfo file faild, %v", err)
	}

	if len(files) > 2 {
		sort.Slice(files, func(i, j int) bool {
			t1, _ := time.Parse("2006-01-02-15-04-05", strings.Split(files[i].Name(), "_")[1])
			t2, _ := time.Parse("2006-01-02-15-04-05", strings.Split(files[j].Name(), "_")[1])
			return t1.Before(t2)
		})
		os.Remove(filepath.Join(path, files[0].Name()))
	}
	return nil
}
