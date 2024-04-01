package product_shelve

import "database/sql"

type ProductShelveStorage struct {
	db *sql.DB
}

func NewProductShelveStorage(db *sql.DB) *ProductShelveStorage {
	return &ProductShelveStorage{db: db}
}
