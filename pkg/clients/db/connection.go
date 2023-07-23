package db

import (
	"companybuilder/config"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type DBInstance struct {
	DB DBInterface
}

// DBInterface ..
type DBInterface interface {
	Query(query string) (*sql.Rows, error)
	Exec(query string, params ...interface{}) (*sql.Result, error)
	GetOne(query string, params ...interface{}) *sql.Row
}

// DB ..
type DBClient struct {
	client *sql.DB
}

func NewConn(appConfig config.IConfig) (*DBInstance, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		appConfig.Get().DB.Host, appConfig.Get().DB.Port, appConfig.Get().DB.User, appConfig.Get().DB.Password, appConfig.Get().DB.Name)

	db, err := sql.Open(config.POSTGRES, psqlInfo)
	if err != nil {
		return nil, errors.New("postgres db client connection failed ->" + err.Error())
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New("postgres db client ping failed ->" + err.Error())
	}

	postgresDbInstance := &DBInstance{
		DB: &DBClient{client: db},
	}

	// Returns
	return postgresDbInstance, nil

}

// Query ..
func (m *DBClient) Query(query string) (*sql.Rows, error) {
	rows, err := m.client.Query(query)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Returns
	return rows, nil
}

func (m *DBClient) Exec(query string, params ...interface{}) (*sql.Result, error) {
	result, err := m.client.Exec(query, params...)
	if err != nil {
		return nil, err
	}

	// Returns
	return &result, nil
}

func (m *DBClient) GetOne(query string, params ...interface{}) *sql.Row {
	return m.client.QueryRow(query, params...)
}
