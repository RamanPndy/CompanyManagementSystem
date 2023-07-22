package config

const (
	POSTGRES      = "postgres"
	COMPANY       = "company"
	GET_ALL_QUERY = "select id, name, description, employees, registered, type, createdAt, updatedAt from %s.%s"
)
