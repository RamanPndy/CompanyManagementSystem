package dal

import (
	"companybuilder/dal/auth"
	"companybuilder/dal/company"
	"companybuilder/shared"
)

// Dal ...
type Dal struct {
	Auth    *auth.Dal
	Company *company.Dal
}

// New creates an instance of Dal object using the dependencies passed
func New(deps *shared.Deps) *Dal {

	// Initialises company dal
	auth := &auth.Dal{
		Deps: deps,
	}

	// Initialises company dal
	company := &company.Dal{
		Deps: deps,
	}

	// Forms dal
	dal := &Dal{
		Auth:    auth,
		Company: company,
	}

	// Returns
	return dal
}
