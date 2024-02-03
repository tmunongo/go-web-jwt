package models

import (
	"github.com/tmunongo/go-web-jwt/pkg/config"

	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

type Toot struct {
	gorm.Model
	content string
	created_at time.Time
	Author string `json:"author"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Toot{})
}

func (t *Toot) NewToot() *Toot {
	db.Create(&t)
	return t
}

func GetAllToots() []Toot {
	var toots []Toot
	db.Find(&toots)
	return toots
}