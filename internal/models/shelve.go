package models

import "database/sql"

type Shelve struct {
	ShelveID int    `json:"shelve_id"`
	Name     string `json:"name"`
}

type ShelveDTO struct {
	ShelveName        string
	ProductName       string
	ProductID         int
	OrderID           int
	Quantity          int
	AdditionalShelves sql.NullString
}
