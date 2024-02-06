package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	database "example.com/go/demoHTTP/database"
	"example.com/go/demoHTTP/web"
	"github.com/go-sql-driver/mysql"
)

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
	store := database.NewSalonStore(db)
	mux := web.NewHandler(store)

	//Démarrer le serveur http
	log.Println("le serveur est lancer sur le port :8097")

	err = http.ListenAndServe(":8097", mux)
	if err != nil {
		log.Fatal(err)
	}

}
