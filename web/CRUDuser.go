package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/go/demoHTTP/models"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) GetUsers() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Je vais sortir un JSON, je rajoute le header correspondant
		writer.Header().Set("Content-Type", "application/json")

		users, _ := h.UserStore.GetUsers()
		err := json.NewEncoder(writer).Encode(users)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handler) AddUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		user := models.User{}
		err := json.NewDecoder(request.Body).Decode(&user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		id, err := h.UserStore.AddUser(user)
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

func (h *Handler) DeleteUser() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		QueryId := chi.URLParam(request, "id")
		id, _ := strconv.Atoi(QueryId)

		err := h.UserStore.DeleteUser(id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(writer, request, "/api", http.StatusSeeOther)
	}
}
