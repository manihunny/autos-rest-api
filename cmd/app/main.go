package main

import (
	"github.com/joho/godotenv"
	"log/slog"
	"main/internal/app"
	"main/internal/config"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file. Error message: ", err)
		os.Exit(1)
	}

	cfg := config.Config{
		ServerHost:  os.Getenv("SERVER_HOST"),
		ServerPort:  os.Getenv("SERVER_PORT"),
		StoragePath: os.Getenv("STORAGE_PATH"),
		DBDriver:    os.Getenv("DB_DRIVER"),
		DBName:      os.Getenv("DB_NAME"),
	}

	app.Run(cfg)
}
