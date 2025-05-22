package models

import "gorm.io/gorm"

type Quote struct {
	gorm.Model
	Quote     string `gorm:"unique;not null;type:varchar(255);default:null"`
	Author    string `gorm:"not null;type:varchar(100);default:null"`
}