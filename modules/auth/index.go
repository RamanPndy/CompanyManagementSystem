package auth

import (
	"companybuilder/dal"
	models "companybuilder/models/auth"
	"companybuilder/shared"
	"context"
)

// AuthInterface ...
type AuthInterface interface {
	Login(context.Context, *models.Request) (*models.Response, error)
	Register(context.Context, *models.Request) (*models.Response, error)
}

// Module..
type Module struct {
	Deps *shared.Deps
	Dal  *dal.Dal
}

// New creates an instance of object using the dependencies passed
func NewModule(deps *shared.Deps) AuthInterface {
	return &Module{
		Deps: deps,
		Dal:  dal.New(deps),
	}
}
