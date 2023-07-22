package auth

import "time"

type User struct {
	ID          int       `db:"id" mapstructure:"id"`
	UserName    string    `db:"username" mapstructure:"name"`
	Password    string    `db:"password" mapstructure:"password"`
	Description *string   `db:"description" mapstructure:"description"`
	IsActive    bool      `db:"isactive" mapstructure:"isactive"`
	CreatedAt   time.Time `db:"createdat" mapstructure:"createdat"`
	UpdatedAt   time.Time `db:"updatedat" mapstructure:"updatedat"`
}
