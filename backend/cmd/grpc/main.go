package main

import (
	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/database"
	grpcservice "github.com/carlosabdoamaral/cbm_brasil/backend/internal/services/grpc_service"
)

func main() {
	common.ReadEnvFile()
	database.Connect()
	grpcservice.Init()
}
