package company

import (
	"companybuilder/dal"
	models "companybuilder/models/company"
	"companybuilder/shared"
	"context"
)

// ConcurrencyInterface ...
type CompanyInterface interface {
	Get(ctx context.Context, id string) (*models.Company, error)
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
