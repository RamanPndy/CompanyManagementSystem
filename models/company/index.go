package company

import "github.com/google/uuid"

type Request struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Employees   int    `json:"employees"`
	Registered  bool   `json:"registered"`
	Type        string `json:"type"`
}

type Company struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Employees   int       `json:"employees"`
	Registered  bool      `json:"registered"`
	Type        string    `json:"type"`
}