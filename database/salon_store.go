package database

import (
	"database/sql"
)

func NewSalonStore(db *sql.DB) *CoiffeurStore {
	return &CoiffeurStore{
		db,
	}
}

type CoiffeurStore struct {
	*sql.DB
}
