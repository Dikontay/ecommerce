package product_order

import (
	"database/sql"
	"ecommerce/internal/models"
)

type ProductOrderStorage struct {
	db *sql.DB
}

func NewProductOrderStorage(db *sql.DB) *ProductOrderStorage {
	return &ProductOrderStorage{db: db}
}

func (s *ProductOrderStorage) GetProductOrder(orderID int) ([]*models.ProductOrder, error) {
	queryProductOrder := `select product_id, order_id, quantity from product_order where order_id = ?`

	rows, err := s.db.Query(queryProductOrder, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var p []*models.ProductOrder
	for rows.Next() {
		var productID, orderID, quantity int

		if err := rows.Scan(&productID, &orderID, &quantity); err != nil {
			return nil, err
		}

		p = append(p, &models.ProductOrder{
			ProductID: productID,
			OrderID:   orderID,
			Quantity:  quantity,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return p, nil
}
