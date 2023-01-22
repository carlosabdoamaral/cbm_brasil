package main

import (
	"fmt"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/api/account"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/api/auth"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/api/occurrence"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/grpc"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/persistence"
	"github.com/gin-gonic/gin"
)

func main() {
	common.ReadEnvVariables()

	_, err := persistence.Connect()
	if err != nil {
		fmt.Printf("[!] %s", err.Error())
		return
	}

	grpc.ConnectToGRPCServer()

	router := gin.Default()
	account.DeclareAccountRoutes(router)
	occurrence.DeclareOccurrenceRoutes(router)
	auth.DeclareAuthRoutes(router)

	router.Run()
}
