package migration

import (
	"product/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1600004192CreateCategoryTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1600004192",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Category{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("categorys").Error
		},
	}
}
