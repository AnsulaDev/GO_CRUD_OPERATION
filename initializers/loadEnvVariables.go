package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// to import env file
func LoadEnvVariables() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
