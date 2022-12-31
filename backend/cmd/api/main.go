package main

import (
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/database"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/handlers"
	grpcservice "github.com/carlosabdoamaral/cbm_brasil/backend/internal/services/grpc_service"
	"github.com/carlosabdoamaral/cbm_brasil/backend/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.ReadEnvFile()
	database.Connect()
	grpcservice.ConnectToGrpcServer()
	InitApi()

}

func InitApi() {
	router := gin.Default()
	router.Use(CORS())

	// TODO: Create sql querys
	tutorialRoutes := router.Group("/tutorials")
	tutorialRoutes.GET("/list", handlers.GetAllTutorials)
	tutorialRoutes.GET("/single", handlers.GetTutorialById)
	tutorialRoutes.POST("/new", handlers.CreateTutorial)
	tutorialRoutes.PUT("/update", handlers.UpdateTutorialById)
	tutorialRoutes.DELETE("/soft-delete", handlers.SoftDeleteTutorialById)
	tutorialRoutes.DELETE("/delete", handlers.DeleteTutorialById)

	occurrenceRoutes := router.Group("/occurrences")
	occurrenceRoutes.GET("/list", handlers.GetAllOccurrences)
	occurrenceRoutes.GET("/by-id", handlers.GetOccurrenceById)
	occurrenceRoutes.GET("/nearby", handlers.GetOccurrencesNearby)
	occurrenceRoutes.POST("/new", handlers.CreateOccurrence)
	occurrenceRoutes.PUT("/update", handlers.UpdateOccurrenceById)
	occurrenceRoutes.DELETE("/soft-delete", handlers.SoftDeleteOccurrenceById)

	// Doing...
	accountRoutes := router.Group("/account")
	accountRoutes.GET("/list", handlers.GetAllAccounts)
	accountRoutes.GET("/private", handlers.GetAccountPrivateDetails)
	accountRoutes.GET("/public", handlers.GetAccountPublicDetails)
	accountRoutes.POST("/new", handlers.CreateAccount)
	accountRoutes.POST("/generate-token", handlers.GenerateAccountToken)
	accountRoutes.PUT("/update", handlers.UpdateAccountById)
	accountRoutes.DELETE("/delete", handlers.DeleteAccountById)
	accountRoutes.DELETE("/soft-delete", handlers.SoftDeleteAccountById)

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
