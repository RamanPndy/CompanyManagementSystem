package company

import (
	"net/http"

	httpUtils "companybuilder/apis/http/utils"
	"companybuilder/modules/company"

	"companybuilder/shared"

	"github.com/gin-gonic/gin"
)

// Company defines the struct which binds an interface to perform
type Company struct {
	CompanyModule company.CompanyInterface
}

func NewCompanyService(deps *shared.Deps) *Company {
	return &Company{
		CompanyModule: company.NewModule(deps),
	}
}

func (c *Company) Get(ctx *gin.Context) {
	var err error

	defer httpUtils.HandleError(ctx, &err, ctx.Request)

	response, err := c.CompanyModule.Get(ctx, ctx.Param("id"))
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, response)
}
