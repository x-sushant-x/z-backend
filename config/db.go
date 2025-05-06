package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// dsn := "user:userpass@tcp(localhost:3306)/taskflow?charset=utf8mb4&parseTime=True&loc=Local"

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
}
