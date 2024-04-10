package product_order

import (
	"database/sql"
	"ecommerce/internal/models"
	"fmt"
	"strings"
)

type ProductOrderStorage struct {
	db *sql.DB
}

type ProductOrderRepo interface {
	GetProductOrderByOrderID(orderID int) (models.ProductOrder, error)
	GetProductOrdersByOrderIDs(orderIDs []int) ([]models.ProductOrder, error)
	GetNumberOfProductOrdersByOrderIDs(orderIDs []int) (int, error)
}

func NewProductOrderStorage(db *sql.DB) *ProductOrderStorage {
	return &ProductOrderStorage{db: db}
}

func (s *ProductOrderStorage) GetProductOrderByOrderID(orderID int) (models.ProductOrder, error) {
	queryProductOrder := `SELECT Product_order_id, order_id, product_id, quantity FROM product_order WHERE order_id = ?`

	row := s.db.QueryRow(queryProductOrder, orderID)

	var id, productID, quantity int

	if err := row.Scan(&id, &orderID, &productID, &quantity); err != nil {
		return models.ProductOrder{}, err
	}

	p := models.ProductOrder{
		ID:        id,
		ProductID: productID,
		OrderID:   orderID,
		Quantity:  quantity,
	}

	if err := row.Err(); err != nil {
		return models.ProductOrder{}, err
	}

	return p, nil
}

func (s *ProductOrderStorage) GetProductOrdersByOrderIDs(orderIDs []int) ([]models.ProductOrder, error) {
	placeholders := strings.Repeat("?,", len(orderIDs)-1) + "?"

	queryProductOrder := fmt.Sprintf(`SELECT Product_order_id, product_id, order_id, quantity FROM product_order WHERE order_id IN (%s)`, placeholders)

	args := make([]interface{}, len(orderIDs))
	for i, id := range orderIDs {
		args[i] = id
	}

	rows, err := s.db.Query(queryProductOrder, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	numberOfProductOrders, err := s.GetNumberOfProductOrdersByOrderIDs(orderIDs)
	if err != nil {
		return nil, err
	}
	productOrders := make([]models.ProductOrder, 0, numberOfProductOrders)
	for rows.Next() {
		var id, productID, orderID, quantity int
		if err := rows.Scan(&id, &productID, &orderID, &quantity); err != nil {
			return nil, err
		}
		productOrders = append(productOrders, models.ProductOrder{
			ID:        id,
			ProductID: productID,
			OrderID:   orderID,
			Quantity:  quantity,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productOrders, nil
}

func (s *ProductOrderStorage) GetNumberOfProductOrdersByOrderIDs(orderIDs []int) (int, error) {
	placeholders := strings.Repeat("?,", len(orderIDs)-1) + "?"

	queryProductOrder := fmt.Sprintf(`SELECT count(*) as number FROM product_order  WHERE order_id IN (%s)`, placeholders)

	args := make([]interface{}, len(orderIDs))
	for i, id := range orderIDs {
		args[i] = id
	}

	row := s.db.QueryRow(queryProductOrder, args...)

	var result int
	if err := row.Scan(&result); err != nil {
		return 0, err
	}
	return result, nil
}
