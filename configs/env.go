package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error loading .env file: %s", err.Error()))
	}
}

func EnvAppPort() string {
	return fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
}

func EnvProjectName() string {
	return os.Getenv("PROJECT_NAME")
}

func EnvMongoURI() string {
	dsn := fmt.Sprintf(
		"mongodb://%s:%s/%s?retryWrites=true&w=majority",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	return dsn
}

func EnvDBUsername() string {
	return os.Getenv("DB_USERNAME")
}

func EnvDBPassword() string {
	return os.Getenv("DB_PASSWORD")
}
