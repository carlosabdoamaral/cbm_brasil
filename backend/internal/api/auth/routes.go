package auth

import "github.com/gin-gonic/gin"

func DeclareAuthRoutes(router *gin.Engine) {
	g := router.Group("/auth")
	g.POST("/sign-in", HandleSignInRequest)
	g.POST("/sign-out", HandleSignOutRequest)
	g.PUT("/recover", HandleRecoverRequest)
}
