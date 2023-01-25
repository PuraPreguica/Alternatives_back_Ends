package models

import "gorm.io/gorm"

type InitialUser struct {
	gorm.Model
	password string `json: "password"`
	email    string `json:"email"`
	fullName string `json: "fullname"`
	user     string `json: "user"`
}

var Users []InitialUser
