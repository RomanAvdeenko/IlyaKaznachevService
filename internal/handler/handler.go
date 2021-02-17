package handler

import (
	"fmt"
	"net/http"
	"workshop/internal/api"
)

type Handler struct {
	client api.Client
}

func NewHandler(client api.Client) *Handler {
	return &Handler{client: client}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	jokeResponse, err := h.client.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, jokeResponse.Joke)
}
