package product_shelf

import (
	"database/sql/driver"
	"ecommerce/internal/models"
	"ecommerce/utils"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

var productShelf = &models.ProductShelf{ShelveID: 3, IsPrimary: true}

func TestGetProductShelfByProductID(t *testing.T) {
	db, mock := utils.NewMock()

	repo := NewProductShelfStorage(db)

	defer db.Close()

	query := `SELECT shelf_id, is_primary FROM product_shelf WHERE product_id = ?`

	rows := sqlmock.NewRows([]string{"shelf_id", "is_primary"}).
		AddRow(productShelf.ShelveID, productShelf.IsPrimary)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(productShelf.ProductID).WillReturnRows(rows)

	result, err := repo.GetProductShelfByProductID(productShelf.ProductID)

	assert.Equal(t, 1, len(result))
	assert.Equal(t, productShelf.ProductID, result[0].ProductID)
	assert.Equal(t, productShelf.ShelveID, result[0].ShelveID)
	assert.Equal(t, productShelf.IsPrimary, result[0].IsPrimary)
	assert.NoError(t, err)
}

func TestGetProductShelvesByProductIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Setup the ProductShelfStorage with the mocked database
	psrepo := NewProductShelfStorage(db)

	// Sample product IDs to search for
	productIDs := []int{1, 2}

	placeholders := strings.Repeat("?,", len(productIDs)-1) + "?"
	query := fmt.Sprintf(`SELECT shelf_id, product_id, isPrimary FROM product_shelf WHERE product_id IN (%s)`, placeholders)
	args := []driver.Value{1, 2}

	rows := sqlmock.NewRows([]string{"shelf_id", "product_id", "isPrimary"}).
		AddRow(1, productIDs[0], true).
		AddRow(2, productIDs[1], false)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(args...).WillReturnRows(rows)

	productShelves, err := psrepo.GetProductShelvesByProductIDs(productIDs)
	if err != nil {
		t.Errorf("error was not expected while getting product shelves: %s", err)
	}

	// Define what we expect to receive
	expected := []models.ProductShelf{
		{ShelveID: 1, ProductID: productIDs[0], IsPrimary: true},
		{ShelveID: 2, ProductID: productIDs[1], IsPrimary: false},
	}

	// Check that the returned slice is as expected
	if len(productShelves) != len(expected) {
		t.Errorf("expected %d product shelves, got %d", len(expected), len(productShelves))
	}
	for i, ps := range productShelves {
		if ps.ShelveID != expected[i].ShelveID || ps.ProductID != expected[i].ProductID || ps.IsPrimary != expected[i].IsPrimary {
			t.Errorf("expected product shelf %v, got %v", expected[i], ps)
		}
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
