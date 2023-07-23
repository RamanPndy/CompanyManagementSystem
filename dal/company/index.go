package company

import (
	"companybuilder/config"
	"companybuilder/dal/common"
	"companybuilder/shared"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

// Dal ...
type Dal struct {
	Deps *shared.Deps
}

func (d *Dal) Get(ctx context.Context) ([]*Company, error) {
	companies := make([]*Company, 0)

	schema, table := common.GetSchemaTable(d.Deps, config.COMPANY)

	query := fmt.Sprintf(config.GET_COMPANIES_QUERY, schema, table)

	rows, err := d.Deps.Database.DB.Query(query)
	if err != nil {
		return companies, err
	}

	for rows.Next() {
		company := &Company{}
		err := rows.Scan(&company.ID, &company.Name, &company.Description, &company.Employees, &company.Registered, &company.Type, &company.CreatedAt, &company.UpdatedAt)
		if err != nil {
			return companies, err
		}
		companies = append(companies, company)
	}

	return companies, nil
}

func (d *Dal) GetOne(ctx context.Context, id string) (*Company, error) {
	schema, table := common.GetSchemaTable(d.Deps, config.COMPANY)

	query := fmt.Sprintf(config.GET_COMPANY_QUERY, schema, table)

	row := d.Deps.Database.DB.GetOne(query, id)

	var company Company
	switch err := row.Scan(&company.ID, &company.Name, &company.Description, &company.Employees, &company.Registered, &company.Type, &company.CreatedAt, &company.UpdatedAt); err {
	case sql.ErrNoRows:
		return nil, errors.New("No company found")
	case nil:
		return &company, nil
	default:
		return nil, err
	}
}

func (d *Dal) Create(ctx context.Context, company *Company) error {
	schema, table := common.GetSchemaTable(d.Deps, config.COMPANY)

	query := fmt.Sprintf(config.CREATE_COMPANY_QUERY, schema, table)

	result, err := d.Deps.Database.DB.Exec(query, company.ID, company.Name, company.Description, company.Employees, company.Registered, company.Type, company.CreatedAt, company.UpdatedAt)
	if err != nil {
		return err
	}

	log.Println(result)

	return nil
}

func (d *Dal) Update(ctx context.Context, company *Company) error {
	schema, table := common.GetSchemaTable(d.Deps, config.COMPANY)

	query := fmt.Sprintf(config.UPDATE_COMPANY_QUERY, schema, table)

	result, err := d.Deps.Database.DB.Exec(query, company.Name, company.Description, company.Employees, company.Registered, company.Type, company.UpdatedAt, company.ID)
	if err != nil {
		return err
	}

	log.Println(result)

	return nil
}

func (d *Dal) Delete(ctx context.Context, id string) error {
	schema, table := common.GetSchemaTable(d.Deps, config.COMPANY)

	query := fmt.Sprintf(config.DELETE_COMPANY_QUERY, schema, table)

	result, err := d.Deps.Database.DB.Exec(query, id)
	if err != nil {
		return err
	}

	log.Println(result)

	return nil
}
