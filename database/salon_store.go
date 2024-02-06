package database

import (
	"database/sql"

	"example.com/go/demoHTTP/models"
)

func NewSalonStore(db *sql.DB) *SalonStore {
	return &SalonStore{
		db,
	}
}

type SalonStore struct {
	*sql.DB
}

func (t *SalonStore) GetSalons() ([]models.Salon, error) {
	var salons []models.Salon

	rows, err := t.Query("SELECT ID, Name, Adresse, Tel FROM Salons")
	if err != nil {
		return []models.Salon{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var salon models.Salon
		if err = rows.Scan(&salon.Adresse, &salon.ID, &salon.Name, &salon.Tel); err != nil {
			return []models.Salon{}, err
		}
		salons = append(salons, salon)
	}

	if err = rows.Err(); err != nil {
		return []models.Salon{}, err
	}

	return salons, nil
}

func (t *SalonStore) AddSalon(item models.Salon) (int, error) {
	// Utilisation de ? comme marqueurs de position pour les valeurs
	res, err := t.DB.Exec("INSERT INTO salons (Name, Tel, Adresse) VALUES (?, ?, ?)", item.Name, item.Tel, item.Adresse)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (t *SalonStore) DeleteSalon(id int) error {
	_, err := t.DB.Exec("DELETE FROM Salon WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (t *SalonStore) ToggleSalon(id int) error {
	_, err := t.DB.Exec("UPDATE Salon SET `completed` = IF (`completed`, 0, 1) WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
