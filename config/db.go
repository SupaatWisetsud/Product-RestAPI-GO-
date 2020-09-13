package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(true)
}

func GetDB() *gorm.DB {
	return db
}
func CloseDB() {
	db.Close()
}
