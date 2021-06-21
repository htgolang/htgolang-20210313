package services

import (
	"cmdb/conf"
	"cmdb/modules"
	"cmdb/utils"
	"log"
	"strconv"
	"strings"
)

func QueryUser(keyWord string) []modules.User {
	users, err := utils.LoadUserInfo(conf.UserInfoPath)
	if err != nil {
		log.Println("load user info faild")
		return nil
	}

	if keyWord == "" {
		return users
	} else {
		tmpuser := []modules.User{}
		for _, v := range users {
			if strings.Contains(strconv.Itoa(v.Id), keyWord) || strings.Contains(v.Name, keyWord) || strings.Contains(v.Addr, keyWord) || strings.Contains(v.Tel, keyWord) {
				tmpuser = append(tmpuser, v)
			}
		}
		return tmpuser
	}

}

func DeleteUser(id int) {
	users, err := utils.LoadUserInfo(conf.UserInfoPath)
	if err != nil {
		log.Println("load user info faild")
		return
	}
	tmpusers := []modules.User{}

	for _, v := range users {
		if v.Id == id {
			continue
		}
		tmpusers = append(tmpusers, v)
	}
	utils.SaveUserInfo(conf.UserInfoPath, tmpusers)
}

func AddUser(username, sex, addr, tel, password, confirmpw string) string {
	users, err := utils.LoadUserInfo(conf.UserInfoPath)
	if err != nil {
		log.Println("load user info faild")
		return ""
	}

	if password != confirmpw {
		return "两次密码不一致"
	}

	if utils.CheckUserName(username, users) {
		return "用户名已存在"
	}

	if utils.CheckTel(tel, users) {
		return "手机号已存在"
	}

	tmpUser := modules.User{Id: utils.GenerateID(users), Name: username, Sex: true, Addr: addr, Tel: tel, Passwd: utils.Md5Encrypt([]byte(password))}
	if sex == "0" {
		tmpUser.Sex = false
	}

	users = append(users, tmpUser)
	utils.SaveUserInfo(conf.UserInfoPath, users)

	return ""
}

func QueryUserByID(id int) *modules.User {
	users, err := utils.LoadUserInfo(conf.UserInfoPath)
	if err != nil {
		log.Println("load user info faild")
		return nil
	}

	for _, v := range users {
		if id == v.Id {
			return &v
		}
	}
	return nil
}

func EditUser(uid int, username, sex, addr, tel string) string {
	users, err := utils.LoadUserInfo(conf.UserInfoPath)
	if err != nil {
		log.Println("load user info faild")
		return ""
	}

	// fmt.Println(users)
	// fmt.Println(uid)
	tmpUsers := []modules.User{}
	user := modules.User{}
	index := 0
	for i, v := range users {
		if uid == v.Id {
			user = v
			index = i
			continue
		}
		tmpUsers = append(tmpUsers, v)

	}

	// fmt.Println(tmpUsers)

	if utils.CheckUserName(username, tmpUsers) {
		return "用户名已存在，请重新输入"
	}

	if utils.CheckTel(tel, tmpUsers) {
		return "号码已存在，请重新输入"
	}

	user.Name = username
	user.Addr = addr
	user.Tel = tel
	if user.Sex && sex == "0" {
		user.Sex = false
	}
	users[index] = user
	utils.SaveUserInfo(conf.UserInfoPath, users)

	return ""

}
