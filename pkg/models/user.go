package models

import (
	"gorm.io/gorm"

	"github.com/tmunongo/go-web-jwt/pkg/config"
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

func GetUserByID(Id int64) (*User, *gorm.DB) {
	var user User

	db := db.Where("id = ?", Id).First(&user)
	return &user, db
}

func GetUserByUsernameAndPassword(username string, password string) (*User, error) {
	var user User

	err := db.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
