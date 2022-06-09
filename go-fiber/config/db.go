package config

import (
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() (err error) {
	DB, err = gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	return
}
