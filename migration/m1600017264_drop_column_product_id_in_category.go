package migration

import (
	"product/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1600017264DropColumnProductIdInCategory() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1600017264",
		Migrate: func(tx *gorm.DB) error {
			return tx.Model(&models.Category{}).DropColumn("product_id").Error
		},
		Rollback: func(tx *gorm.DB) error {
			type Category struct {
				ProductID uint
			}
			return tx.AutoMigrate(&Category{}).Error
		},
	}
}
