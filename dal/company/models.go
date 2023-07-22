package company

import "database/sql"

type Company struct {
	ID          string         `db:"id" mapstructure:"id"`
	Name        sql.NullString `db:"name" mapstructure:"name"`
	Description sql.NullString `db:"description" mapstructure:"description"`
	Employees   sql.NullInt32  `db:"employees" mapstructure:"employees"`
	Registered  sql.NullBool   `db:"registered" mapstructure:"registered"`
	Type        sql.NullString `db:"type" mapstructure:"type"`
	CreatedAt   sql.NullTime   `db:"createdAt" mapstructure:"createdAt"`
	UpdatedAt   sql.NullTime   `db:"updatedAt" mapstructure:"updatedAt"`
}
