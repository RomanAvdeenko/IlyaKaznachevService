package main

import (
	"log"
	"net/http"

	"github.com/RomanAvdeenko/IlyaKaznachevService/internal/handler"
	"github.com/go-chi/chi"
)

func main() {
	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("Can't start server")
	}

}
