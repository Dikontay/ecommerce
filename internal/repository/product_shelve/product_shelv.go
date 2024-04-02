package product_shelve

import (
	"database/sql"
	"ecommerce/internal/models"
)

type ProductShelveStorage struct {
	db *sql.DB
}

func NewProductShelveStorage(db *sql.DB) *ProductShelveStorage {
	return &ProductShelveStorage{db: db}
}

func (s *ProductShelveStorage) GetProductShelve(productID int) ([]*models.ProductShelve, error) {
	queryProductOrder := `select  shelve_id, isPrimary from product_shelve where product_id = ?`

	rows, err := s.db.Query(queryProductOrder, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var p []*models.ProductShelve
	for rows.Next() {
		var shelveID int
		var isPrimary bool

		if err := rows.Scan(&shelveID, &isPrimary); err != nil {
			return nil, err
		}

		p = append(p, &models.ProductShelve{
			ShelveID:  shelveID,
			IsPrimary: isPrimary,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return p, nil
}

func (s *ProductShelveStorage) GetProductShelves(productIDs []int) ([]*models.ProductShelve, error) {
	queryProductOrder := `select  shelve_id, isPrimary from product_shelve where product_id = ?`

	rows, err := s.db.Query(queryProductOrder, productIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var p []*models.ProductShelve
	for rows.Next() {
		var shelveID int
		var isPrimary bool

		if err := rows.Scan(&shelveID, &isPrimary); err != nil {
			return nil, err
		}

		p = append(p, &models.ProductShelve{
			ShelveID:  shelveID,
			IsPrimary: isPrimary,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return p, nil
}
