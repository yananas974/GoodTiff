package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/go/demoHTTP/models"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) GetSalons() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Je vais sortir un JSON, je rajoute le header correspondant
		writer.Header().Set("Content-Type", "application/json")

		salons, _ := h.SalonStore.GetSalons()
		err := json.NewEncoder(writer).Encode(salons)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) AddSalon() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		salon := models.Salon{}
		err := json.NewDecoder(request.Body).Decode(&salon)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := h.SalonStore.AddSalon(salon)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(writer).Encode(struct {
			Status     string `json:"status"`
			Message    string `json:"message"`
			NewSalonId int    `json:"newTodoId"`
		}{
			Status:     "success",
			Message:    "Nouveau salon inséré avec succès",
			NewSalonId: id,
		})
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) DeleteSalon() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)

		err := h.SalonStore.DeleteSalon(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(writer, request, "/api", http.StatusSeeOther)
	}
}

func (h *Handler) ToggleSalon() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)

		err := h.SalonStore.ToggleSalon(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(writer, request, "/api", http.StatusSeeOther)
	}
}
