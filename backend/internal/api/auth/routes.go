package auth

import "github.com/gin-gonic/gin"

func DeclareAuthRoutes(router *gin.Engine) {
	g := router.Group("/auth")
	g.POST("/sign-in/cpf", HandleSignInRequestByCPF)
	g.POST("/sign-in/email", HandleSignInRequestByEmail)
	g.POST("/sign-out", HandleSignOutRequest)
	g.PUT("/recover", HandleRecoverRequest)
}
