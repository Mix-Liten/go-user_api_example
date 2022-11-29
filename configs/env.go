package configs

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func ENVAPPPort() string {
	return fmt.Sprintf(":%s", os.Getenv("PROJECT_NAME"))
}

func ENVProjectName() string {
	return fmt.Sprintf(":%s", os.Getenv("PROJECT_NAME"))
}

func EnvMongoURI() string {
	dsn := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s?retryWrites=true&w=majority",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	return dsn
}
