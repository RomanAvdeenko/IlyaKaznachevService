package handler

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return new(Handler)
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}
