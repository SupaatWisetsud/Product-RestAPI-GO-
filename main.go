package main

import (
	"log"
	"os"
	"product/config"
	"product/migration"
	"product/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	config.InitDB()
	migration.Migraiton()

	uploadFile := []string{"users", "products"}
	for _, dir := range uploadFile {
		os.MkdirAll("uploads/"+dir, 0755)
	}

	r := gin.Default()

	r.Static("/uploads", "uploads")

	routes.Serve(r)

	r.Run(":" + os.Getenv("PORT"))
}
