package product_shelf

import (
	"database/sql"
	"ecommerce/internal/models"
	"fmt"
	"strings"
)

type ProductShelfStorage struct {
	db *sql.DB
}

type ProductShelfRepo interface {
	GetProductShelve(productID int) ([]models.ProductShelve, error)
	GetProductShelvesByProductIDs(productIDs []int) ([]models.ProductShelve, error)
}

func NewProductShelfStorage(db *sql.DB) *ProductShelfStorage {
	return &ProductShelfStorage{db: db}
}

func (s *ProductShelfStorage) GetProductShelve(productID int) ([]models.ProductShelve, error) {
	queryProductOrder := `select  shelve_id, isPrimary from product_shelve where product_id = ?`

	rows, err := s.db.Query(queryProductOrder, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var p []models.ProductShelve
	for rows.Next() {
		var shelveID int
		var isPrimary bool

		if err := rows.Scan(&shelveID, &isPrimary); err != nil {
			return nil, err
		}

		p = append(p, models.ProductShelve{
			ShelveID:  shelveID,
			IsPrimary: isPrimary,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return p, nil
}

func (s *ProductShelfStorage) GetProductShelvesByProductIDs(productIDs []int) ([]models.ProductShelve, error) {
	placeholders := strings.Repeat("?,", len(productIDs)-1) + "?"

	queryProductShelve := fmt.Sprintf(`SELECT shelve_id, product_id,  isPrimary FROM product_shelve WHERE product_id IN (%s)`, placeholders)

	args := make([]interface{}, len(productIDs))
	for i, id := range productIDs {
		args[i] = id
	}

	rows, err := s.db.Query(queryProductShelve, args...)
	if err != nil {
		return nil, err
	}

	var productShelves []models.ProductShelve
	for rows.Next() {
		var shelveID, productID int
		var isPrimary bool

		if err := rows.Scan(&shelveID, &productID, &isPrimary); err != nil {
			return nil, err
		}

		productShelves = append(productShelves, models.ProductShelve{
			ShelveID:  shelveID,
			IsPrimary: isPrimary,
			ProductID: productID,
		})
	}

	defer rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productShelves, nil
}
