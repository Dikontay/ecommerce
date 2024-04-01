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

func NewProductOrderStorage(db *sql.DB) *ProductOrderStorage {
	return &ProductOrderStorage{db: db}
}

func (s *ProductOrderStorage) GetOrdersInfo(orderIDs []int) ([]*models.ShelveDTO, error) {

	placeholders := strings.Repeat("?,", len(orderIDs)-1) + "?"

	query := fmt.Sprintf(`
	SELECT
	  s.name AS shelve_name,
	  p.name AS product_name,
	  p.ID AS product_id,
	  po.order_id,
	  po.quantity,
	  (
		SELECT GROUP_CONCAT(ps2.shelve_id)
		FROM product_shelve ps2
		WHERE ps2.product_id = p.ID AND ps2.isPrimary = 0
	  ) AS additional_shelves
	FROM
	  product_order po
	JOIN products p ON po.product_id = p.ID
	JOIN product_shelve ps ON p.ID = ps.product_id AND ps.isPrimary = 1
	JOIN shelves s ON ps.shelve_id = s.shelve_id
	WHERE
	  po.order_id IN (%s)
	ORDER BY
	  s.name, p.name;
	`, placeholders)

	args := make([]interface{}, len(orderIDs))
	for i, id := range orderIDs {
		args[i] = id
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []*models.ShelveDTO
	for rows.Next() {
		var shelveName, productName string
		var additionalShelves sql.NullString
		var productID, orderID, quantity int
		if err := rows.Scan(&shelveName, &productName, &productID, &orderID, &quantity, &additionalShelves); err != nil {
			return nil, err
		}
		details = append(details, &models.ShelveDTO{
			ShelveName:        shelveName,
			ProductName:       productName,
			ProductID:         productID,
			OrderID:           orderID,
			Quantity:          quantity,
			AdditionalShelves: additionalShelves,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return details, nil
}
