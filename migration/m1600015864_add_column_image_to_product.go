package migration

import (
	"product/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1600015864AddColumnImageToProduct() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1600015864",
		Migrate: func(tx *gorm.DB) error {
			type Product struct {
				Image string `gorm:"not null"`
			}
			return tx.AutoMigrate(&Product{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Model(&models.Product{}).DropColumn("image").Error
		},
	}
}
