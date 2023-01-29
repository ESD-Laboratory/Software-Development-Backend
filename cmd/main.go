package main

import (
	"os"

	"github.com/ESD-Laboratory/Software-Development-Backend/internal/config"
	"github.com/ESD-Laboratory/Software-Development-Backend/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error Loading .env filer")
	}

	config.Connect(&config.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	})

	server.Start()
}
