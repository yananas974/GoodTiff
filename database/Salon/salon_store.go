package database

import (
	"database/sql"
	"fmt"

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

	rows, err := t.Query("SELECT id, name, adresse, tel FROM Salons")
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête SQL:", err)
		return []models.Salon{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var salon models.Salon
		if err = rows.Scan(&salon.ID, &salon.Name, &salon.Adresse, &salon.Tel); err != nil {
			fmt.Println("Erreur lors de la lecture des lignes:", err)
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
	res, err := t.DB.Exec("INSERT INTO Salons (Name, Tel, Adresse) VALUES ( ?, ?, ?)")
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête SQL:", err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (t *SalonStore) DeleteSalon(id int) error {
	_, err := t.DB.Exec("DELETE FROM Salons WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (t *SalonStore) ToggleSalon(id int) error {
	_, err := t.DB.Exec("UPDATE Salons SET `completed` = IF (`completed`, 0, 1) WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
