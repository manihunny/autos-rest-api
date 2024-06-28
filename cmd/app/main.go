package main

import (
	"github.com/joho/godotenv"
	"log"
	"main/internal/app"
	"main/internal/config"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Error message: ", err.Error())
	}

	cfg := config.Config{
		ServerHost: os.Getenv("SERVER_HOST"),
		ServerPort: os.Getenv("SERVER_PORT"),
		DBDriver:   os.Getenv("DB_DRIVER"),
		DBName:     os.Getenv("DB_NAME"),
	}

	app.Run(cfg)
}
