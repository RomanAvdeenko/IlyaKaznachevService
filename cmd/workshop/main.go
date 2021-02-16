package main

import (
	"log"
	"net/http"

	"github.com/RomanAvdeenko/IlyaKaznachevService/internal/config"
	"github.com/RomanAvdeenko/IlyaKaznachevService/internal/handler"
	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)
	err = http.ListenAndServe(cfg.Host+":"+cfg.Port, r)
	if err != nil {
		log.Fatalln("Can't start server")
	}

}
