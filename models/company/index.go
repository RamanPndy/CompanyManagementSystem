package company

import (
	"time"
)

type CreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Employees   int    `json:"employees" binding:"required"`
	Registered  bool   `json:"registered" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

type UpdateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Employees   int    `json:"employees"`
	Registered  *bool  `json:"registered"`
	Type        string `json:"type"`
}

type Company struct {
	ID          string    `json:"id" mapstructure:"id"`
	Name        string    `json:"name" mapstructure:"name"`
	Description string    `json:"description" mapstructure:"description"`
	Employees   int       `json:"employees" mapstructure:"employees"`
	Registered  bool      `json:"registered" mapstructure:"registered"`
	Type        string    `json:"type" mapstructure:"type"`
	CreatedAt   time.Time `json:"createdat" mapstructure:"createdat"`
	UpdatedAt   time.Time `json:"updatedat" mapstructure:"updatedat"`
}

type Response struct {
	Id      *string `json:"id,omitempty"`
	Message string  `json:"message"`
}
