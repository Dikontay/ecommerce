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
	GetProductShelfByProductID(productID int) ([]models.ProductShelf, error)
	GetProductShelvesByProductIDs(productIDs []int) ([]models.ProductShelf, error)
	GetNumberOfProductShelfByProductIDs(productIDs []int) (int, error)
}

func NewProductShelfStorage(db *sql.DB) *ProductShelfStorage {
	return &ProductShelfStorage{db: db}
}

func (s *ProductShelfStorage) GetProductShelfByProductID(productID int) ([]models.ProductShelf, error) {
	query := `SELECT shelf_id, is_primary FROM product_shelf WHERE product_id = ?`

	rows, err := s.db.Query(query, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productShelves []models.ProductShelf
	for rows.Next() {
		var ps models.ProductShelf
		// Only scan the fields you need, assuming ProductID is already known
		if err := rows.Scan(&ps.ShelveID, &ps.IsPrimary); err != nil {
			return nil, err
		}
		ps.ProductID = productID // Set the ProductID since it's known
		productShelves = append(productShelves, ps)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productShelves, nil
}

func (s *ProductShelfStorage) GetProductShelvesByProductIDs(productIDs []int) ([]models.ProductShelf, error) {
	placeholders := strings.Repeat("?,", len(productIDs)-1) + "?"

	queryProductShelve := fmt.Sprintf(`SELECT shelf_id, product_id,  isPrimary FROM product_shelf WHERE product_id IN (%s)`, placeholders)

	args := make([]interface{}, len(productIDs))
	for i, id := range productIDs {
		args[i] = id
	}

	rows, err := s.db.Query(queryProductShelve, args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	numberOfProductShelf, err := s.GetNumberOfProductShelfByProductIDs(productIDs)
	if err != nil {
		return nil, err
	}

	productShelves := make([]models.ProductShelf, 0, numberOfProductShelf)
	for rows.Next() {
		var shelveID, productID int
		var isPrimary bool

		if err := rows.Scan(&shelveID, &productID, &isPrimary); err != nil {
			return nil, err
		}

		productShelves = append(productShelves, models.ProductShelf{
			ShelveID:  shelveID,
			IsPrimary: isPrimary,
			ProductID: productID,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productShelves, nil
}

func (s *ProductShelfStorage) GetNumberOfProductShelfByProductIDs(productIDs []int) (int, error) {
	placeholders := strings.Repeat("?,", len(productIDs)-1) + "?"

	queryProductOrder := fmt.Sprintf(`SELECT count(*) as number FROM product_shelf  WHERE product_id IN (%s)`, placeholders)

	args := make([]interface{}, len(productIDs))
	for i, id := range productIDs {
		args[i] = id
	}

	row := s.db.QueryRow(queryProductOrder, args...)

	var result int
	if err := row.Scan(&result); err != nil {
		return 0, err
	}

	return result, nil
}
