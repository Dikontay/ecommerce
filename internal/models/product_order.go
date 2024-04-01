package models

type ProductOrder struct {
	ProductOrderID int     `json:"product_order_id"`
	ProductID      int     `json:"product_id"`
	OrderID        int     `json:"order_id"`
	Quantity       int     `json:"quantity"`
	Discount       float64 `json:"discount"`
}

type ProductOrderDto struct {
}
