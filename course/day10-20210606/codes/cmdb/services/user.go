package services

import "cmdb/models"

var users = []models.User{
	models.User{1, "kk", true},
	models.User{2, "arune", true},
	models.User{3, "fangping", true},
	models.User{4, "wangxian", true},
	models.User{5, "lanhaibin", true},
	models.User{6, "ada", false},
}

func QueryUsers(q string) []models.User {
	return users
}

func DeleteUser(id int) {
	tempUsers := make([]models.User, 0, len(users))
	for _, user := range users {
		if user.ID == id {
			continue
		}
		tempUsers = append(tempUsers, user)
	}
	users = tempUsers
}
