package account

import "github.com/gin-gonic/gin"

func DeclareAccountRoutes(router *gin.Engine) {
	g := router.Group("/account")
	g.POST("/new", HandleNewAccountRequest)
	g.POST("/by-id", HandleFetchAccountByIdRequest)
	g.PUT("/edit", HandleEditAccountRequest)
	g.DELETE("/soft-delete", HandleSoftDeleteAccountRequest)
	g.PUT("/recover", HandleRecoverRequest)
}
