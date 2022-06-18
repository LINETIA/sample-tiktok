package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(32);not null"`
	Nickname string `gorm:"type:varchar(32);not null"`
	Password string `gorm:"type:varchar(100);not null"`
}
