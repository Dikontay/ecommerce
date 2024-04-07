package shelf

import (
	"ecommerce/internal/models"
	"ecommerce/utils"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

var shelf = &models.Shelf{
	ShelfID: 1,
	Name:    "A",
}

func TestGetShelfByID(t *testing.T) {
	db, mock := utils.NewMock()

	repo := NewShelfStorage(db)

	defer db.Close()

	query := `SELECT * FROM shelves WHERE shelf_id = ?`

	rows := sqlmock.NewRows([]string{"shelf_id", "name"}).
		AddRow(shelf.ShelfID, shelf.Name)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(shelf.ShelfID).WillReturnRows(rows)
	result, err := repo.GetShelfByID(shelf.ShelfID)

	assert.Equal(t, shelf.ShelfID, result.ShelfID)
	assert.Equal(t, shelf.Name, result.Name)

	assert.NoError(t, err)

}
