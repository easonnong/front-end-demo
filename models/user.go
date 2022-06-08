package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"username"`
	Password string `json:"password"`
}
