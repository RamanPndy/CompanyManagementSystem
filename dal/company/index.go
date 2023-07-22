package company

import (
	"companybuilder/config"
	"companybuilder/shared"
	"context"
	"fmt"
)

// Dal ...
type Dal struct {
	Deps *shared.Deps
}

func (d *Dal) getSchemaTable(tableName string) (string, string) {
	dbTables := d.Deps.Config.Get().DB.Tables
	for _, dbTable := range dbTables {
		if dbTable.Name == tableName {
			return dbTable.Schema, dbTable.Name
		}
	}

	return "", ""
}

func (d *Dal) GetCompanies(ctx context.Context) ([]*Company, error) {
	companies := make([]*Company, 0)

	schema, table := d.getSchemaTable(config.COMPANY)

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
