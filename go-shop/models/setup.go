package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost user=postgres dbname=go_shop sslmode=disable password=yourpassword")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB.AutoMigrate(&User{})
}
