package handlers

import "cmdb/models"

type dataFormat struct {
	Ok   string        `json:"ok,omitempty"`
	Msg  string        `json:"msg,omitempty"`
	Data []models.User `json:"data,omitempty"`
}
