package db

import (
	"companybuilder/config"
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DBInstance struct {
	DB DBInterface
}

// DBInterface ..
type DBInterface interface {
	Query(query string) (*sql.Rows, error)
}

// DB ..
type DBClient struct {
	client *sql.DB
}

func NewConn(appConfig config.IConfig) (*DBInstance, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PWD")
	dbName := os.Getenv("DB_NAME")

	if dbHost != "" && dbPort != "" && dbUser != "" && dbPwd != "" && dbName != "" {
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPwd, dbName)

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
	return nil, errors.New("DB credentials not present")
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
