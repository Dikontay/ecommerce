package models

import "database/sql"

type Shelf struct {
	ShelfID int    `json:"shelve_id"`
	Name    string `json:"name"`
}

type ShelveDTO struct {
	ShelveName        string
	ProductName       string
	ProductID         int
	OrderID           int
	Quantity          int
	AdditionalShelves sql.NullString
}
