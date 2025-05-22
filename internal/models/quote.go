package models

import "gorm.io/gorm"

type Quote struct {
	gorm.Model
	Quote     string `gorm:"not null"`
	Author    string `gorm:"not null"`
}