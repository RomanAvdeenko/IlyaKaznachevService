package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"workshop/internal/api/jokes"
	"workshop/internal/config"
	"workshop/internal/handler"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatalln(err)
	}

	apiClient := jokes.NewJokeClient(cfg.JokeURL)

	h := handler.NewHandler(apiClient, cfg.CustomJoke)
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	serverAddr := cfg.Host + ":" + cfg.Port

	srv := &http.Server{
		Addr:    serverAddr,
		Handler: r,
	}

	// handle server shutdown garsefully
	quit := make(chan os.Signal)
	done := make(chan error)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		//
		// ... Make shutdown jobs
		//
		done <- err
	}()
	//

	log.Println("Starting server at ", serverAddr)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalln("Can't start server.", err)
	}

	err = <-done

	shutingMsg := "Shuting down server..."
	if err != nil {
		shutingMsg = fmt.Sprintf("%s with error %v\n", shutingMsg, err)
	}
	log.Print(shutingMsg)
}
