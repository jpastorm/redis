package bootstrap

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func loadConfiguration() (string, string, error){
	err := godotenv.Load()
	if err != nil {
		return "","", errors.New("Error loading .env file")
	}
	token := os.Getenv("TOKEN")
	api := os.Getenv("API")

	return token, api, nil
}
