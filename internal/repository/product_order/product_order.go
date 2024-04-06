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
	GetProductOrderByID(orderID int) (models.ProductOrder, error)
	GetProductOrdersByIDs(orderIDs []int) ([]models.ProductOrder, error)
}

func NewProductOrderStorage(db *sql.DB) *ProductOrderStorage {
	return &ProductOrderStorage{db: db}
}

func (s *ProductOrderStorage) GetProductOrderByID(orderID int) (models.ProductOrder, error) {
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

func (s *ProductOrderStorage) GetProductOrdersByIDs(orderIDs []int) ([]models.ProductOrder, error) {
	placeholders := strings.Repeat("?,", len(orderIDs)-1) + "?"
	fmt.Println(placeholders)

	queryProductOrder := fmt.Sprintf(`SELECT Product_order_id, product_id, order_id, quantity FROM product_order WHERE order_id IN (%s)`, placeholders)

	// Подготавливаем аргументы для запроса
	args := make([]interface{}, len(orderIDs))
	for i, id := range orderIDs {
		args[i] = id
	}

	// Теперь используем args... для передачи аргументов
	rows, err := s.db.Query(queryProductOrder, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var productOrders []models.ProductOrder
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
