package models

type Order struct {
	ID           int    `json:"id"`
	OrderDate    string `json:"order_date"`    // In a real-world application, consider using time.Time
	ShippingDate string `json:"shipping_date"` // In a real-world application, consider using time.Time
}

type OrderInfoDTO struct {
	ProductName string
	ProductID   int
	OrderID     int
	Quantity    int
}

type OrderPageDTO struct {
	ShelveName string
	Orders     []*OrderInfoDTO
}
