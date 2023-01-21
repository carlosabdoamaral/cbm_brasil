package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariables() error {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		return errEnv
	}

	DatabaseUser = os.Getenv("DB_USER")
	DatabasePass = os.Getenv("DB_PASS")
	DatabaseHost = os.Getenv("DB_HOST")
	DatabaseName = os.Getenv("DB_NAME")
	DatabasePort = os.Getenv("DB_PORT")
	DatabaseSSL = os.Getenv("DB_SSL")
	DatabaseDriver = os.Getenv("DB_DRIVER")

	RabbitURL = os.Getenv("RABBIT_URL")
	RabbitAccountQueueName = os.Getenv("RABBIT_ACCOUNT_SERVICE_QUEUENAME")
	RabbitOccurrenceQueueName = os.Getenv("RABBIT_OCCURRENCE_SERVICE_QUEUENAME")

	return nil
}

func ReadEnvVariables() {
	envErr := GetEnvVariables()
	if envErr != nil {
		fmt.Println(envErr.Error())
		fmt.Println("Failed to load env file. Make sure .env file exists!")
	}
}
