package product_order

import (
	"database/sql/driver"
	"ecommerce/internal/models"
	"ecommerce/utils"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var productOrder = &models.ProductOrder{
	ID:        2,
	OrderID:   1,
	ProductID: 2,
	Quantity:  3,
}

func TestGetProductOrderByID(t *testing.T) {
	db, mock := utils.NewMock()

	repo := NewProductOrderStorage(db)

	defer db.Close()

	query := `SELECT Product_order_id, order_id, product_id, quantity FROM product_order WHERE order_id = ?`

	rows := sqlmock.NewRows([]string{"Product_order_id", "product_id", "order_id", "quantity"}).
		AddRow(productOrder.ID, productOrder.OrderID, productOrder.ProductID, productOrder.Quantity)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(productOrder.OrderID).WillReturnRows(rows)
	p, err := repo.GetProductOrderByOrderID(productOrder.OrderID)
	assert.NotNil(t, p)
	assert.NoError(t, err)
}

func TestGetProductOrdersByIDs(t *testing.T) {
	db, mock := utils.NewMock()

	repo := NewProductOrderStorage(db)
	defer db.Close()

	orderIDs := []int{1, 2}

	placeholders := strings.Repeat("?,", len(orderIDs)-1) + "?"
	rows := sqlmock.NewRows([]string{"Product_order_id", "product_id,", "order_id", "quantity"}).
		AddRow(1, 1, 2, 5).
		AddRow(2, 2, 2, 4)

	args := []driver.Value{1, 2}
	queryProductOrder := fmt.Sprintf(`SELECT Product_order_id, product_id, order_id, quantity FROM product_order WHERE order_id IN (%s)`, placeholders)

	mock.ExpectQuery(regexp.QuoteMeta(queryProductOrder)).WithArgs(args...).WillReturnRows(rows)

	productOrders, err := repo.GetProductOrdersByOrderIDs(orderIDs)
	if err != nil {
		t.Errorf("error was not expected while getting data from product_order table %s", err)
	}

	expected := []models.ProductOrder{{
		ID:        1,
		ProductID: 1,
		OrderID:   2,
		Quantity:  5,
	}, {
		ID:        2,
		ProductID: 2,
		OrderID:   2,
		Quantity:  4,
	}}

	if !reflect.DeepEqual(productOrders, expected) {
		t.Errorf("expected %#v, got %#v", expected, productOrders)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations: %s", err)
	}
}
