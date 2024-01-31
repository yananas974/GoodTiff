package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/go/demoHTTP/models"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) GetSalon() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Je vais sortir un JSON, je rajoute le header correspondant
		writer.Header().Set("Content-Type", "application/json")

	}
}

func (h *Handler) AddSalon() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != "POST" {
			http.Error(writer, "cette route n'est disponible qu'en POST", http.StatusBadRequest)
			return
		}

		item := models.Salon{}
		err := json.NewDecoder(request.Body).Decode(&item)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		h.Salons = append(h.Salons, item)

		// Je crée une struct anonyme à laquelle j'ajoute
		// tout de suite un contenu que je renvoie en JSON
		err = json.NewEncoder(writer).Encode(struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "Nouveau todo inséré avec succès",
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) DeleteSalon() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Ce routeur me permet d'accéder facilement aux paramètres
		QueryId := chi.URLParam(request, "id")
		// Attention, ce sont toujours des strings...
		id, _ := strconv.Atoi(QueryId)

		for index, todo := range h.Salons {
			if id == todo.ID {
				h.Salons = append(h.Salons[:index], h.Salons[index+1:]...)
				break
			}
		}

		http.Redirect(writer, request, "/", http.StatusSeeOther)
	}
}
