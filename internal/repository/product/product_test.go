package product

import (
	"database/sql"
	"ecommerce/internal/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

var product = &models.Product{
	ID:    1,
	Name:  "Ipad",
	Price: 289.9,
}

func TestGetProductByID(t *testing.T) {
	db, mock := NewMock()

	repo := NewProductStorage(db)

	defer db.Close()

	query := "SELECT ID, name, price FROM products WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"ID", "name", "price"}).
		AddRow(product.ID, product.Name, product.Price)

	mock.ExpectQuery(query).WithArgs(product.ID).WillReturnRows(rows)

	p, err := repo.GetProductByID(product.ID)
	assert.NotNil(t, p)
	assert.NoError(t, err)

}

func TestGetProductsByIDs(t *testing.T) {

}
