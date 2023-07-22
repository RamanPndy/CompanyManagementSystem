package dal

import (
	"companybuilder/dal/company"
	"companybuilder/shared"
)

// Dal ...
type Dal struct {
	Company *company.Dal
}

// New creates an instance of Dal object using the dependencies passed
func New(deps *shared.Deps) *Dal {

	// Initialises company dal
	company := &company.Dal{
		Deps: deps,
	}

	// Forms dal
	dal := &Dal{
		Company: company,
	}

	// Returns
	return dal
}
