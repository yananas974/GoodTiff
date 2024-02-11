package web

import (
	salonDB "example.com/go/demoHTTP/database/salon"
	userDB "example.com/go/demoHTTP/database/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Stores struct {
	SalonStore *salonDB.SalonStore
	UserStore  *userDB.UserStore
}

func NewHandler(stores *Stores) *Handler {
	handler := &Handler{
		Mux:        chi.NewRouter(),
		SalonStore: stores.SalonStore,
		UserStore:  stores.UserStore,
	}

	handler.Use(middleware.Logger)

	handler.Route("/salons", func(r chi.Router) {
		r.Get("/", handler.GetSalons())
		r.Post("/add", handler.AddSalon())
		r.Delete("/{id}", handler.DeleteSalon())
		
	})

	handler.Route("/user", func(r chi.Router) {
		r.Get("/", handler.GetUsers())
		r.Post("/add", handler.AddUser())
		r.Delete("/{id}", handler.DeleteUser())
	})

	return handler
}

type Handler struct {
	*chi.Mux
	*salonDB.SalonStore
	*userDB.UserStore
}
