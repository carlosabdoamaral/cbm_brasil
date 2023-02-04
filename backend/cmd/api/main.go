package main

import (
	"fmt"

	"github.com/carlosabdoamaral/cbm_brasil/backend/common"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/api/account"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/api/auth"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/api/occurrence"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/grpc"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/persistence"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/rabbit"
	"github.com/gin-gonic/gin"
)

func main() {
	common.GetEnvVariables()

	_, err := persistence.Connect()
	if err != nil {
		fmt.Printf("[!] %s", err.Error())
		return
	}

	rabbit.Connect()
	grpc.ConnectToGRPCServer()

	router := gin.Default()
	router.Use(CORS())
	account.DeclareAccountRoutes(router)
	occurrence.DeclareOccurrenceRoutes(router)
	auth.DeclareAuthRoutes(router)

	router.Run()
}

// https://github.com/gin-contrib/cors/issues/29
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
