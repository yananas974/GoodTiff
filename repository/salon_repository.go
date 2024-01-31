package repository

import (
	"database/sql"

	"example.com/go/demoHTTP/models"
)

type SalonRepository struct {
	db *sql.DB
}

func NewSalon(db *sql.DB) *SalonRepository {
	return &SalonRepository{
		db: db,
	}
}

func (r *SalonRepository) GetAllSalons() ([]models.Salon, error) {
	query := "SELECT * FROM salon"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	salons := []models.Salon{}
	for rows.Next() {
		salon := models.Salon{}
		err := rows.Scan(&salon.ID, &salon.Name, salon.Tel, salon.Adresse)

		if err != nil {
			return nil, err
		}
		salons = append(salons, salon)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return salons, err
}
