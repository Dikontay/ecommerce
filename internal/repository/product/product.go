package product

import (
	"database/sql"
	"ecommerce/internal/models"
)

type ProductStorage struct {
	db *sql.DB
}

type ProductRepo interface {
	GetProductByID(id int) (models.Product, error)
}

func NewProductStorage(db *sql.DB) *ProductStorage {
	return &ProductStorage{db: db}
}

func (db *ProductStorage) GetProductByID(id int) (models.Product, error) {
	query := `SELECT * from products where ID = ?`

	row := db.db.QueryRow(query, id)

	product := models.Product{}
	err := row.Scan(&product.ID,
		&product.Name,
		&product.Price,
	)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
