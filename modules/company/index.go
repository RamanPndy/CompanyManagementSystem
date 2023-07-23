package company

import (
	"companybuilder/dal"
	models "companybuilder/models/company"
	"companybuilder/shared"
	"context"
)

// CompanyInterface ...
type CompanyInterface interface {
	GetAll(context.Context) ([]*models.Company, error)
	Get(context.Context, string) (*models.Company, error)
	Create(context.Context, *models.CreateRequest) (*models.Response, error)
	Update(context.Context, string, *models.UpdateRequest) (*models.Response, error)
	Delete(context.Context, string) (*models.Response, error)
}

// Module..
type Module struct {
	Deps *shared.Deps
	Dal  *dal.Dal
}

// New creates an instance of object using the dependencies passed
func NewModule(deps *shared.Deps) CompanyInterface {
	return &Module{
		Deps: deps,
		Dal:  dal.New(deps),
	}
}
