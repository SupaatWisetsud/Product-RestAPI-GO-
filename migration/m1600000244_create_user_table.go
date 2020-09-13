package migration

import (
	"product/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1600000244CreateUserTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1600000244",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.User{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("users").Error
		},
	}
}
