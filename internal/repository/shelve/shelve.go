package shelve

import (
	"database/sql"
	"ecommerce/internal/models"
)

type ShelveStorage struct {
	db *sql.DB
}

func NewShelveStorage(db *sql.DB) *ShelveStorage {
	return &ShelveStorage{db: db}
}

func (s *ShelveStorage) GetShelvesByIDs(shelveIDs []int) ([]*models.Shelve, error) {
	query := `select * from shelves where shelve_id = ?`
	rows, err := s.db.Query(query, shelveIDs)
	if err != nil {
		return nil, err
	}
	var shelves []*models.Shelve
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, err
		}

		shelves = append(shelves, &models.Shelve{
			ShelveID: id,
			Name:     name,
		})
	}

	return shelves, nil
}
