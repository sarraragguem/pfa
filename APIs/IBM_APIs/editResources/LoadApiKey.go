package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetApiKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apikey := os.Getenv("API_KEY")
	return apikey
}
