package config

const (
	POSTGRES          = "postgres"
	COMPANY           = "company"
	USERS             = "users"
	GET_ALL_QUERY     = "select id, name, description, employees, registered, type, createdat, updatedat from %s.%s"
	CREATE_USER_QUERY = `INSERT INTO %s.%s (username, password, isactive, createdat, updatedat) VALUES ($1, $2, $3, $4, $5)`
	GET_USER_QUERY    = "select id, username, password, description, isactive, createdat, updatedat from %s.%s WHERE username =$1"
)
