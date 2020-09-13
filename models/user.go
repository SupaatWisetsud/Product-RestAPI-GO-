package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Avater   string
	Role     string `gorm:"default:'Member';not null"`
}

func (u *User) GenerateHashPassword() string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	return string(hash)
}
