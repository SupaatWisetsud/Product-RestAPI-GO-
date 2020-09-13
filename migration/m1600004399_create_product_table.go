package migration

import (
	"product/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1600004399CreateProductTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1600004399",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Product{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("products").Error
		},
	}
}
