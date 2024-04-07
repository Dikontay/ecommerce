package models

type Order struct {
	ID           int    `json:"id"`
	OrderDate    string `json:"order_date"`
	ShippingDate string `json:"shipping_date"`
}

type OrderInfoDTO struct {
	Product      Product      //{id, name}
	ProductOrder ProductOrder //{order_id, quantity}
}

type OrderPageDTO struct {
	Shelve *Shelf //{id, shelve name }
	Orders []*OrderInfoDTO
}
