package shelf

import (
	"database/sql"
	"ecommerce/internal/models"
)

type ShelfStorage struct {
	db *sql.DB
}

func NewShelfStorage(db *sql.DB) *ShelfStorage {
	return &ShelfStorage{db: db}
}

type ShelfRepo interface {
	GetShelfByID(shelfID int) (*models.Shelf, error)
}

func (s *ShelfStorage) GetShelfByID(shelfID int) (*models.Shelf, error) {
	query := `SELECT * FROM shelves WHERE shelf_id = ?`
	row := s.db.QueryRow(query, shelfID)

	var id int
	var name string
	if err := row.Scan(&id, &name); err != nil {
		return nil, err
	}

	shelf := &models.Shelf{
		ShelfID: id,
		Name:    name,
	}

	return shelf, nil
}
