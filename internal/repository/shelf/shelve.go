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
	GetShelveByID(shelveID int) (*models.Shelf, error)
	GetShelvesByIDs(shelveIDs []int) ([]*models.Shelf, error)
}

func (s *ShelfStorage) GetShelveByID(shelveID int) (*models.Shelf, error) {
	query := `select * from shelves where shelve_id = ?`
	row := s.db.QueryRow(query, shelveID)

	var id int
	var name string
	if err := row.Scan(&id, &name); err != nil {
		return nil, err
	}

	shelve := &models.Shelf{
		ShelfID: id,
		Name:    name,
	}

	return shelve, nil
}

func (s *ShelfStorage) GetShelvesByIDs(shelveIDs []int) ([]*models.Shelf, error) {
	query := `select * from shelves where shelve_id = ?`
	rows, err := s.db.Query(query, shelveIDs)
	if err != nil {
		return nil, err
	}
	var shelves []*models.Shelf
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		shelves = append(shelves, &models.Shelf{
			ShelfID: id,
			Name:    name,
		})
	}

	return shelves, nil
}
