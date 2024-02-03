package models

import (
	"gorm.io/gorm"

	"github.com/tmunongo/go-web-jwt/config"
)

// var db *gorm.DB

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) NewUser() *User {
	db.Create(&u)
	return u
}