package main

import (
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/database"
	grpcservice "github.com/carlosabdoamaral/cbm_brasil/backend/internal/services/grpc_service"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/utils"
)

func main() {
	utils.ReadEnvFile()
	database.Connect()
	grpcservice.Init()
}
