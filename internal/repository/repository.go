package repository

import (
	"database/sql"
	"ecommerce/internal/repository/order"
	"ecommerce/internal/repository/product"
	"ecommerce/internal/repository/product_order"
	"ecommerce/internal/repository/product_shelve"
	"ecommerce/internal/repository/shelve"
)

type Repository struct {
	Order         order.OrderStorage
	Product       product.ProductStorage
	Shelve        shelve.ShelveStorage
	ProductShelve product_shelve.ProductShelveStorage
	ProductOrder  product_order.ProductOrderStorage
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Order:         *order.NewOrderStorage(db),
		Product:       *product.NewProductStorage(db),
		Shelve:        *shelve.NewShelveStorage(db),
		ProductShelve: *product_shelve.NewProductShelveStorage(db),
		ProductOrder:  *product_order.NewProductOrderStorage(db),
	}
}
