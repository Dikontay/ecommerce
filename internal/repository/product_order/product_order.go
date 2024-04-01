package product_order

import "database/sql"

type ProductOrderStorage struct {
	db *sql.DB
}

func NewProductOrderStorage(db *sql.DB) *ProductOrderStorage {
	return &ProductOrderStorage{db: db}
}
