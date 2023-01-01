package main

import (
	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/database"
	rabbitservice "github.com/carlosabdoamaral/cbm_brasil/backend/internal/services/rabbit_service"
)

func main() {
	common.ReadEnvFile()
	database.Connect()
	rabbitservice.Connect()
	rabbitservice.DeclareQueue()
	rabbitservice.Consumer()
}
