package occurrence

import "github.com/gin-gonic/gin"

func DeclareOccurrenceRoutes(router *gin.Engine) {
	g := router.Group("/occurrence")
	g.POST("/new", HandleNewOccurrenceRequest)
	g.GET("/by-id", HandleFetchOccurrenceByIdRequest)
	g.GET("/all", HandleFetchAllOccurrencesRequest)
	g.GET("/nearby", HandleFetchNearbyOccurrencesRequest)
	g.PUT("/edit", HandleEditOccurrenceRequest)
	g.PUT("/accept", HandleAcceptOccurrenceRequest)
	g.PUT("/cancel", HandleCancelOccurrenceRequest)
	g.PUT("/finish", HandleFinishOccurenceRequest)
	g.DELETE("/soft-delete", HandleSoftDeleteOccurrenceById)
}
