package company

import (
	"companybuilder/config"
	"companybuilder/dal/company"
	models "companybuilder/models/company"
	"companybuilder/pkg/utils"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

func (m *Module) GetAll(ctx context.Context) ([]*models.Company, error) {
	result, err := m.Dal.Company.Get(ctx)
	if err != nil {
		return nil, err
	}

	var output []*models.Company
	for _, log := range result {
		c := models.Company{}
		err := mapstructure.Decode(log, &c)
		if err != nil {
			return nil, err
		}
		c.CreatedAt = log.CreatedAt
		c.UpdatedAt = log.UpdatedAt
		output = append(output, &c)
	}

	return output, nil
}

func (m *Module) Get(ctx context.Context, id string) (*models.Company, error) {
	result, err := m.Dal.Company.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}

	var company models.Company
	err = mapstructure.Decode(result, &company)
	if err != nil {
		return nil, err
	}

	company.CreatedAt = result.CreatedAt
	company.UpdatedAt = result.UpdatedAt

	return &company, nil
}

func (m *Module) Create(ctx context.Context, request *models.CreateRequest) (*models.Response, error) {
	if !utils.Contains(strings.Split(config.COMPANY_TYPES, "|"), request.Type) {
		return nil, errors.New("Invalid Company Type")
	}

	company := company.Company{
		ID:          uuid.NewString(),
		Name:        request.Name,
		Description: request.Description,
		Employees:   request.Employees,
		Registered:  request.Registered,
		Type:        request.Type,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := m.Dal.Company.Create(ctx, &company)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Id:      &company.ID,
		Message: "Company " + request.Name + " Created Successfully",
	}, nil
}

func (m *Module) Update(ctx context.Context, id string, request *models.UpdateRequest) (*models.Response, error) {
	company, err := m.Dal.Company.GetOne(ctx, id)
	if err != nil {
		return nil, err
	}

	if request.Name != "" {
		company.Name = request.Name
	}

	if request.Description != "" {
		company.Description = request.Description
	}

	if request.Employees != 0 {
		company.Employees = request.Employees
	}

	if request.Registered != nil {
		company.Registered = *request.Registered
	}

	if request.Type != "" {
		company.Type = request.Type
	}

	company.UpdatedAt = time.Now()

	err = m.Dal.Company.Update(ctx, company)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Id:      &company.ID,
		Message: "Company " + id + " Updated Successfully",
	}, nil
}

func (m *Module) Delete(ctx context.Context, id string) (*models.Response, error) {
	err := m.Dal.Company.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Message: "Company " + id + " Deleted Successfully",
	}, nil

}
