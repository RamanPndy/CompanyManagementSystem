package auth

import (
	"companybuilder/shared"

	"github.com/gin-gonic/gin"
)

// NewRoute Creates and initializes route
func NewRoute(router *gin.Engine, deps *shared.Deps) {
	bindRoutes(router, deps)
}

func bindRoutes(router *gin.Engine, deps *shared.Deps) {
	authService := NewAuthService(deps)
	routerAPI := router.Group("/auth")
	{
		routerAPI.POST("/register", authService.Register)
		routerAPI.POST("/login", authService.Login)
	}
}
