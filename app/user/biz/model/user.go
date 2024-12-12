package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	email          string `gorm:"uniqueIndex"`
	PasswordHashed string
}

func (User) TableName() string {
	return "user"
}
