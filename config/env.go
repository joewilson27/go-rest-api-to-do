package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvir(key string) string {
	return os.Getenv(key)
}

func LoadEnvir() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		fmt.Println("No App Env")
		env = "development"
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
}
