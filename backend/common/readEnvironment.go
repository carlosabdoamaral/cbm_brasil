package common

import (
	"os"

	"github.com/joho/godotenv"
)

func ReadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		LogError(err.Error())
		return
	}

	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PORT = os.Getenv("DB_PORT")
	DB_SSL = os.Getenv("DB_SSL")
	DB_DRIVER = os.Getenv("DB_DRIVER")
	RABBIT_URL = os.Getenv("RABBIT_URL")
	OCCURRENCE_QUEUE = os.Getenv("OCCURRENCE_QUEUE")
	MOCK_MODE = os.Getenv("MOCK_MODE")
}
