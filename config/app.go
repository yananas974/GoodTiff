package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"example.com/go/demoHTTP/models"
	"github.com/go-sql-driver/mysql"
)

var Todos = []models.Salon{
	{
		ID:      1,
		Name:    "Faire la vaisselle",
		Tel:     0606060606,
		Adresse: "35 rue ",
	},
}

//var db *sql.DB

func Connection() {

	conf := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "coiffeur",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected")

	//config du router
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bienvenue sur la page "))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("bienvenue sur la page "))
	})

	//Démarrer le serveur http
	log.Println("le serveur est lancer sur le port :8097")

	err = http.ListenAndServe(":8097", mux)
	if err != nil {
		log.Fatal(err)
	}

}
