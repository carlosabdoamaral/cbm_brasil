package main

import (
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/database"
	rabbitservice "github.com/carlosabdoamaral/cbm_brasil/backend/internal/services/rabbit_service"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/utils"
)

func main() {
	utils.ReadEnvFile()
	database.Connect()
	rabbitservice.Connect()
	rabbitservice.DeclareQueue()
	rabbitservice.Consumer()
}
