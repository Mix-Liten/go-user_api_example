package helpers

import (
	"log"
	"os"
)

func verifyEnv() {
	keys := []string{
		"APP_PORT",
		"JWT_SECRET",
		"JWT_TTL",
	}
	for _, key := range keys {
		if os.Getenv(key) == "" {
			log.Fatalf("env %s is required", key)
		}
	}
}
