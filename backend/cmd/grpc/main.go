package main

import (
	"fmt"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/grpc"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/persistence"
)

func main() {
	common.GetEnvVariables()

	_, err := persistence.Connect()
	if err != nil {
		fmt.Printf("[!] %s", err.Error())
		return
	}

	grpc.ServeGrpcServer()
}
