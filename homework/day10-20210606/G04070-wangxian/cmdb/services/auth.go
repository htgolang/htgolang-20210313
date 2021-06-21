package services

import (
	"cmdb/conf"
	"cmdb/modules"
	"cmdb/utils"
	"log"
)

func Login(username, password string) *modules.User {
	users, err := utils.LoadUserInfo(conf.UserInfoPath)
	if err != nil {
		log.Println("load user info faild")
		return nil
	}

	for _, user := range users {
		if username == user.Name && utils.Md5Encrypt([]byte(password)) == user.Passwd {
			return &user
		}
	}
	return nil
}
