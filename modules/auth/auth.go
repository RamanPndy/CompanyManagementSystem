package auth

import (
	"context"
	"errors"
	"time"

	dalModels "companybuilder/dal/auth"
	models "companybuilder/models/auth"
	"companybuilder/pkg/utils"
)

func (m *Module) Register(ctx context.Context, request *models.CreateRequest) (*models.Response, error) {
	escapedUsername := utils.EscapeString(request.Username)
	hashedPassword, err := utils.CreateHash(request.Password)
	if err != nil {
		return nil, err
	}

	user := &dalModels.User{
		UserName:  escapedUsername,
		Password:  hashedPassword,
		IsActive:  request.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if request.Description != "" {
		user.Description = &request.Description
	}

	err = m.Dal.Auth.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Message: utils.GetPointerToString("User Created Successfully"),
	}, nil
}

func (m *Module) Login(ctx context.Context, request *models.LoginRequest) (*models.Response, error) {
	user, err := m.Dal.Auth.GetUser(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	err = utils.VerifyHash(request.Password, user.Password)
	if err != nil {
		return nil, errors.New("Password is wrong for the user: " + request.Username)
	}

	if !user.IsActive {
		return nil, errors.New("User " + user.UserName + "is not active")
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Token: &token,
	}, nil
}

func (m *Module) Update(ctx context.Context, request *models.UpdateRequest) (*models.Response, error) {
	user, err := m.Dal.Auth.GetUser(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	if request.Password != "" {
		user.Password, err = utils.CreateHash(request.Password)
		if err != nil {
			return nil, err
		}
	}

	if request.Description != "" {
		user.Description = &request.Description
	}

	if request.IsActive != nil {
		user.IsActive = *request.IsActive
	}

	user.UpdatedAt = time.Now()

	err = m.Dal.Auth.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return &models.Response{
		Message: utils.GetPointerToString("User updated Successfully"),
	}, nil
}
