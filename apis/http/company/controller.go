package company

import (
	"net/http"

	httpUtils "companybuilder/apis/http/utils"
	models "companybuilder/models/company"
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

func (c *Company) GetAll(ctx *gin.Context) {
	var err error

	defer httpUtils.HandleError(ctx, &err, ctx.Request)

	response, err := c.CompanyModule.GetAll(ctx)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Company) Create(ctx *gin.Context) {
	var request *models.CreateRequest
	var err error

	defer httpUtils.HandleError(ctx, &err, ctx.Request)

	// Binds Request
	if err = ctx.BindJSON(&request); err != nil {
		return
	}

	response, err := c.CompanyModule.Create(ctx, request)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Company) Update(ctx *gin.Context) {
	var request *models.UpdateRequest
	var err error

	defer httpUtils.HandleError(ctx, &err, ctx.Request)

	// Binds Request
	if err = ctx.BindJSON(&request); err != nil {
		return
	}

	response, err := c.CompanyModule.Update(ctx, ctx.Param("id"), request)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Company) Delete(ctx *gin.Context) {
	var err error
	defer httpUtils.HandleError(ctx, &err, ctx.Request)

	response, err := c.CompanyModule.Delete(ctx, ctx.Param("id"))
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, response)
}
