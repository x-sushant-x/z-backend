package config

import (
	"fmt"
	"github.com/x-sushant-x/Zocket/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	once.Do(func() {
		dsn := os.Getenv("DSN")
		if len(dsn) == 0 {
			log.Fatal("DSN not provided in env.")
		}

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to DB: ", err)
		}

		DB = db
		fmt.Println("Database connected!")
	})
}

func AutoMigrateDB() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Task{})
}
