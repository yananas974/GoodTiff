package store

import (
	"database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "demo.db")
	if err != nil {
		log.Fatal(err)
	}

	// Créer la table 'salons' si elle n'existe pas
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS salons (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			Name TEXT,
			Tel INTEGER,
			Adresse TEXT
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Créer la table 'users' si elle n'existe pas
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			Name TEXT,
			Lastname TEXT,
			Tel INTEGER,
			Email TEXT,
			Password TEXT
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}
