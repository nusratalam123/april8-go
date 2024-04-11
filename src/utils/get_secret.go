package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetSecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	return os.Getenv("MONGO_URI")
}
