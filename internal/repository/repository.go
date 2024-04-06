package repository

import (
	"database/sql"
	"ecommerce/internal/repository/order"
	"ecommerce/internal/repository/product"
	"ecommerce/internal/repository/product_order"
	"ecommerce/internal/repository/product_shelf"
	"ecommerce/internal/repository/shelf"
)

type Repository struct {
	Order         order.OrderRepo
	Product       product.ProductRepo
	Shelve        shelf.ShelfRepo
	ProductShelve product_shelf.ProductShelfRepo
	ProductOrder  product_order.ProductOrderRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Order:         order.NewOrderStorage(db),
		Product:       product.NewProductStorage(db),
		Shelve:        shelf.NewShelfStorage(db),
		ProductShelve: product_shelf.NewProductShelfStorage(db),
		ProductOrder:  product_order.NewProductOrderStorage(db),
	}
}
