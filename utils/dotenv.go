package util

import (
	"os"

	"github.com/joho/godotenv"
)

func GodotEnv(key string) string {
	godotenv.Load(".env")

	return os.Getenv(key)
}
