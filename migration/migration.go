package migration

import (
	"log"
	"product/config"

	"gopkg.in/gormigrate.v1"
)

func Migraiton() {
	db := config.GetDB()
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		m1600000244CreateUserTable(),
		m1600004192CreateCategoryTable(),
		m1600004399CreateProductTable(),
		m1600015864AddColumnImageToProduct(),
		m1600017264DropColumnProductIdInCategory(),
	})

	if err := m.Migrate(); err != nil {
		log.Fatal(err)
	}
}
