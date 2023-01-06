package config

import (
	"fmt"
	"os"

	"github.com/athun/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("\nerror while connectiong the databases..")
	}

	DB.AutoMigrate(&models.User{})
	return DB
}
