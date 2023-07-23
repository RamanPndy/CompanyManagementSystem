package auth

import (
	"net/http"

	httpUtils "companybuilder/apis/http/utils"
	models "companybuilder/models/auth"
	"companybuilder/modules/auth"

	"companybuilder/shared"

	"github.com/gin-gonic/gin"
)

// Company defines the struct which binds an interface to perform
type Auth struct {
	AuthModule auth.AuthInterface
}

func NewAuthService(deps *shared.Deps) *Auth {
	return &Auth{
		AuthModule: auth.NewModule(deps),
	}
}

func (a *Auth) Register(ctx *gin.Context) {
	var request *models.CreateRequest
	var err error

	defer httpUtils.HandleError(ctx, &err, ctx.Request)

	// Binds Request
	if err = ctx.BindJSON(&request); err != nil {
		return
	}

	response, err := a.AuthModule.Register(ctx, request)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (a *Auth) Login(ctx *gin.Context) {
	var request *models.LoginRequest
	var err error

	defer httpUtils.HandleError(ctx, &err, ctx.Request)

	// Binds Request
	if err = ctx.BindJSON(&request); err != nil {
		return
	}

	response, err := a.AuthModule.Login(ctx, request)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (a *Auth) Update(ctx *gin.Context) {
	var request *models.UpdateRequest
	var err error

	defer httpUtils.HandleError(ctx, &err, ctx.Request)

	// Binds Request
	if err = ctx.BindJSON(&request); err != nil {
		return
	}

	response, err := a.AuthModule.Update(ctx, request)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, response)
}
