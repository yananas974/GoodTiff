package database

import (
	"database/sql"
	"fmt"

	"example.com/go/demoHTTP/models"
)

func NewUseStore(db *sql.DB) *UserStore {
	return &UserStore{
		db,
	}
}

type UserStore struct {
	*sql.DB
}

func (t *UserStore) GetUsers() ([]models.User, error) {
	var users []models.User

	rows, err := t.Query("SELECT id, name, lastname, tel, email, password FROM User")
	if err != nil {
		fmt.Println("Erreur lors de l'exécution de la requête SQL:", err)
		return []models.User{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Lastname, &user.Tel, &user.Email, &user.Password); err != nil {
			fmt.Println("Erreur lors de la lecture des lignes:", err)
			return []models.User{}, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return []models.User{}, err
	}

	return users, nil
}

func (t *UserStore) AddUser(item models.User) (int, error) {
	// Utilisation de ? comme marqueurs de position pour les valeurs
	res, err := t.DB.Exec("INSERT INTO User (Name, Lastname, Tel, Email, Password) VALUES ( ?, ?, ?, ?, ?)")
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

func (t *UserStore) DeleteUser(id int) error {
	_, err := t.DB.Exec("DELETE FROM User WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
