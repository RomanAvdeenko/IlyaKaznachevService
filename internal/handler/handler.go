package handler

import (
	"fmt"
	"net/http"
	"workshop/internal/api"
)

// Handler is a API handler
type Handler struct {
	client api.Client
}

// NewHandler is a Handler creare helper
func NewHandler(client api.Client) *Handler {
	return &Handler{client: client}
}

// Hello is a handler func for Joke API
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	jokeResponse, err := h.client.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, jokeResponse.Joke)
}
