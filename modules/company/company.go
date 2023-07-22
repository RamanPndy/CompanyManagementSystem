package company

import (
	models "companybuilder/models/company"
	"context"
)

func (m *Module) Get(ctx context.Context, id string) (*models.Company, error) {
	return &models.Company{}, nil
}
