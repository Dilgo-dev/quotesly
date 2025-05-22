package config

import (
	"github.com/Dilgo-dev/quotesly/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Quote{}); err != nil {
		return nil, err
	}

	DB = db

	return DB, nil
}

func GetDB() *gorm.DB {
	return DB
}