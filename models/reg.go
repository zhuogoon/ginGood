package models

import "gorm.io/gorm"

type RequestBody struct {
	gorm.Model
	Username string `json:"username" binding:"required" gorm:"column:username;type:varchar(255);not null;uniqueIndex:unique_index_key"`
	Password string `json:"password" binding:"required" gorm:"column:password;type:varchar(255);not null;uniqueIndex:unique_index_key"`
}

type ResponseBody struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
