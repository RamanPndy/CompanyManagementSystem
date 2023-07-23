package company

import (
	"time"
)

type Company struct {
	ID          string    `db:"id" mapstructure:"id"`
	Name        string    `db:"name" mapstructure:"name"`
	Description string    `db:"description" mapstructure:"description"`
	Employees   int       `db:"employees" mapstructure:"employees"`
	Registered  bool      `db:"registered" mapstructure:"registered"`
	Type        string    `db:"type" mapstructure:"type"`
	CreatedAt   time.Time `db:"createdat" mapstructure:"createdat"`
	UpdatedAt   time.Time `db:"updatedat" mapstructure:"updatedat"`
}
