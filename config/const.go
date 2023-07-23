package config

const (
	DEVELOPMENT          = "development"
	DOCKER               = "docker"
	POSTGRES             = "postgres"
	COMPANY              = "company"
	USERS                = "users"
	GET_COMPANIES_QUERY  = "select id, name, description, employees, registered, type, createdat, updatedat from %s.%s"
	GET_COMPANY_QUERY    = "select id, name, description, employees, registered, type, createdat, updatedat from %s.%s WHERE id =$1"
	CREATE_COMPANY_QUERY = "INSERT INTO %s.%s (id, name, description, employees, registered, type, createdat, updatedat) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	UPDATE_COMPANY_QUERY = "UPDATE %s.%s SET name = $1, description = $2, employees = $3, registered = $4, type = $5, updatedat = $6 WHERE id = $7"
	DELETE_COMPANY_QUERY = "DELETE from %s.%s WHERE id =$1"
	CREATE_USER_QUERY    = "INSERT INTO %s.%s (username, password, isactive, createdat, updatedat) VALUES ($1, $2, $3, $4, $5)"
	GET_USER_QUERY       = "select id, username, password, description, isactive, createdat, updatedat from %s.%s WHERE username =$1"
)
