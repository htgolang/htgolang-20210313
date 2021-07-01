package services

import (
	"cmdb/models"
	"cmdb/utils"
)

type dataFormat struct {
	Ok   string        `json:"ok,omitempty"`
	Msg  string        `json:"msg,omitempty"`
	Data []models.User `json:"data,omitempty"`
}

var users = []models.User{
	models.User{1, "yizuo", utils.Md5sum("yizuo"), true, "Wuhan", "66666666"},
	models.User{2, "saber", utils.Md5sum("saber"), false, "Wuhan", "66666666"},
	models.User{3, "leimu", utils.Md5sum("leimu"), false, "Shanghai", "88888888"},
}
