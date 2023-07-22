package auth

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

func (d *Dal) CreateUser(ctx context.Context, user *User) error {
	schema, table := common.GetSchemaTable(d.Deps, config.USERS)

	query := fmt.Sprintf(config.CREATE_USER_QUERY, schema, table)

	result, err := d.Deps.Database.DB.Exec(query, user.UserName, user.Password, user.IsActive, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	log.Println(result)

	return nil
}

func (d *Dal) GetUser(ctx context.Context, username string) (*User, error) {
	schema, table := common.GetSchemaTable(d.Deps, config.USERS)

	query := fmt.Sprintf(config.GET_USER_QUERY, schema, table)

	row := d.Deps.Database.DB.GetOne(query, username)

	var user User
	switch err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Description, &user.IsActive, &user.CreatedAt, &user.UpdatedAt); err {
	case sql.ErrNoRows:
		return nil, errors.New("No user found")
	case nil:
		return &user, nil
	default:
		return nil, err
	}
}
