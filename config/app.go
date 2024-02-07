package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	salon "example.com/go/demoHTTP/database/salon"
	user "example.com/go/demoHTTP/database/user"

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

	store := salon.NewSalonStore(db)
	userStore := user.NewUserStore(db)

	stores := &web.Stores{
		SalonStore: store,
		UserStore:  userStore,
	}

	mux := web.NewHandler(stores)

	//Démarrer le serveur http
	log.Println("le serveur est lancer sur le port :8097")

	err = http.ListenAndServe(":8097", mux)
	if err != nil {
		log.Fatal(err)
	}

}
