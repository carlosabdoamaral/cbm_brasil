package occurrence

import "github.com/gin-gonic/gin"

func HandleNewOccurrenceRequest(ctx *gin.Context)          {}
func HandleFetchOccurrenceByIdRequest(ctx *gin.Context)    {}
func HandleFetchAllOccurrencesRequest(ctx *gin.Context)    {}
func HandleFetchNearbyOccurrencesRequest(ctx *gin.Context) {}
func HandleEditOccurrenceRequest(ctx *gin.Context)         {}
func HandleAcceptOccurrenceRequest(ctx *gin.Context)       {}
func HandleCancelOccurrenceRequest(ctx *gin.Context)       {}
func HandleFinishOccurenceRequest(ctx *gin.Context)        {}
func HandleSoftDeleteOccurrenceById(ctx *gin.Context)      {}
