package web

import (
	database "example.com/go/demoHTTP/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler(store *database.SalonStore) *Handler {
	handler := &Handler{
		chi.NewRouter(),
		store,
	}

	handler.Use(middleware.Logger)

	handler.Route("/api", func(r chi.Router) {
		r.Get("/", handler.GetSalons())
		r.Post("/", handler.AddSalon())
		r.Delete("/{id}", handler.DeleteSalon())
		r.Patch("/{id}", handler.ToggleSalon())
	})

	return handler
}

type Handler struct {
	*chi.Mux
	*database.SalonStore
}
