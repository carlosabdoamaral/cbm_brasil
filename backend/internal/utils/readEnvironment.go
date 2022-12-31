package utils

import (
	"os"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/joho/godotenv"
)

func ReadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		common.LogError(err.Error())
		return
	}

	common.DB_USER = os.Getenv("DB_USER")
	common.DB_PASS = os.Getenv("DB_PASS")
	common.DB_HOST = os.Getenv("DB_HOST")
	common.DB_NAME = os.Getenv("DB_NAME")
	common.DB_PORT = os.Getenv("DB_PORT")
	common.DB_SSL = os.Getenv("DB_SSL")
	common.DB_DRIVER = os.Getenv("DB_DRIVER")
	common.RABBIT_URL = os.Getenv("RABBIT_URL")
	common.OCCURRENCE_QUEUE = os.Getenv("OCCURRENCE_QUEUE")
	common.MOCK_MODE = os.Getenv("MOCK_MODE")
}
