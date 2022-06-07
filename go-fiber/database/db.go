package database

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)
