package services

import (
	"cmdb/models"
)

func ValidateLogin(name string, password string) *models.User {
	if name == "kk" && password == "123456" {
		return &models.User{1, "kk", true}
	}
	return nil
}
