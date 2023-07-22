package company

import (
	"companybuilder/shared"

	"github.com/gin-gonic/gin"
)

// NewRoute Creates and initializes route
func NewRoute(router *gin.Engine, deps *shared.Deps) {
	bindRoutes(router, deps)
}

func bindRoutes(router *gin.Engine, deps *shared.Deps) {
	companyService := NewCompanyService(deps)
	routerAPI := router.Group("/company")
	{
		// routerAPI.GET("/", companyService.GetAll)
		routerAPI.GET("/:id", companyService.Get)
		// routerAPI.POST("/", companyService.Create)
		// routerAPI.PATCH("/:id", companyService.Update)
		// routerAPI.DELETE("/:id", companyService.Delete)
	}
}
