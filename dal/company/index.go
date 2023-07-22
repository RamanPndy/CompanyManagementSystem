package company

import (
	"companybuilder/config"
	"companybuilder/dal/common"
	"companybuilder/shared"
	"context"
	"fmt"
)

// Dal ...
type Dal struct {
	Deps *shared.Deps
}

func (d *Dal) GetCompanies(ctx context.Context) ([]*Company, error) {
	companies := make([]*Company, 0)

	schema, table := common.GetSchemaTable(d.Deps, config.COMPANY)

	query := fmt.Sprintf(config.GET_ALL_QUERY, schema, table)

	rows, err := d.Deps.Database.DB.Query(query)
	if err != nil {
		return companies, err
	}

	for rows.Next() {
		company := &Company{}
		err := rows.Scan(company)
		if err != nil {
			return companies, err
		}
		companies = append(companies, company)
	}

	return companies, nil
}
