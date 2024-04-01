package shelve

import "database/sql"

type ShelveStorage struct {
	db *sql.DB
}

func NewShelveStorage(db *sql.DB) *ShelveStorage {
	return &ShelveStorage{db: db}
}
