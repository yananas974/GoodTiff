package web

import (
	"fmt"
	"net/http"
	"example.com/go/demoHTTP/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


type Handler struct {
	*chi.Mux
}
func NewHandler(salons []models.Salon) *Handler {
	handler := &Handler{
		chi.NewRouter(),
		salons,
	}

	type Handler struct {
		*chi.Mux
		Salons []models.Salon
	}


	// Utiliser le middleware Logger
	handler.Use(middleware.Logger)

	// Configurer les routes avec différentes méthodes
	handler.Get("/", handler.GetSalon)
	handler.Post("/", handler.AddSalon)
	handler.Delete("/{id}", handler.DeleteSalon)

	return handler
}



// Fonction pour gérer la méthode GET
func (h *Handler) GetSalon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bienvenue sur la page. (GET)\n"))
}

// Fonction pour gérer la méthode POST
func (h *Handler) AddSalon(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ajouter un salon. (POST)\n"))
	// Implémenter la logique pour ajouter un salon si nécessaire
}

// Fonction pour gérer la méthode DELETE
func (h *Handler) DeleteSalon(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("Supprimer le salon avec l'ID %s. (DELETE)\n", id)))
	// Implémenter la logique pour supprimer un salon si nécessaire
}

