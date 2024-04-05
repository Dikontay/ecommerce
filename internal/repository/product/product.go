package product

import (
	"database/sql"
	"ecommerce/internal/models"
)

type ProductStorage struct {
	db *sql.DB
}

type ProductRepo interface {
	GetProductByID(id int) (*models.Product, error)
	GetAllProducts() (*[]models.Product, error)
}

func NewProductStorage(db *sql.DB) *ProductStorage {
	return &ProductStorage{db: db}
}

func (db *ProductStorage) GetProductByID(id int) (*models.Product, error) {
	query := `SELECT * from products where ID = ?`

	row := db.db.QueryRow(query, id)

	product := &models.Product{}
	err := row.Scan(&product.ID,
		&product.Name,
		&product.Price,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (db *ProductStorage) GetProductsByIDs(productIDs []int) ([]*models.Product, error) {
	query := `select * from products where ID in (?)`

	products := make([]*models.Product, 0)
	rows, err := db.db.Query(query, productIDs)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := &models.Product{}

		err = rows.Scan(&product.ID,
			&product.Name,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}
	}

	return products, nil
}
