package database

import (
	"database/sql"

	"example.com/go/demoHTTP/models"
)

func CreateStore(db *sql.DB) *Store {
	return &Store{
		NewSalonStore(db),
	}
}

type Store struct {
	models.SalonStoreInterface
}
