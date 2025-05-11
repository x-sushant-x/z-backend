package main

import (
	"github.com/joho/godotenv"
	"github.com/x-sushant-x/Zocket/api"
	"github.com/x-sushant-x/Zocket/config"
)

func init() {
	godotenv.Load()
}

func main() {
	config.ConnectDB()
	config.AutoMigrateDB()

	config.NewOpenAIClient()
	config.NewGeminiClient()

	api.StartServer()
}
