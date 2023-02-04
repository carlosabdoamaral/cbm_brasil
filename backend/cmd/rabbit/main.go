package main

import (
	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/grpc"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/rabbit"
)

func main() {
	common.GetEnvVariables()

	grpc.ConnectToGRPCServer()

	rabbit.Connect()
	rabbit.DeclareQueue()
	rabbit.StartConsuming()
}
