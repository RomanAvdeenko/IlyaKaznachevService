package handler

import (
	"fmt"
	"net/http"
	"workshop/internal/api"
)

// Handler is a API handler
type Handler struct {
	client     api.Client
	customJoke string
}

// NewHandler is a Handler creare helper
func NewHandler(client api.Client, customJoke string) *Handler {
	return &Handler{
		client:     client,
		customJoke: customJoke,
	}
}

// Hello is a handler func for Joke API
func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	if h.customJoke != "" {
		fmt.Fprint(w, h.customJoke)
		return
	}
	jokeResponse, err := h.client.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, jokeResponse.Joke)
}
