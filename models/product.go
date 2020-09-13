package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name       string `gorm:"unique;not null"`
	Price      uint   `gorm:"not null"`
	Image      string `gorm:"not null"`
	CategoryID uint
	Category   Category
}
