package main

import (
	"fmt"
	"github.com/3110Y/idsaver/internal/infrastructure/di"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dependencyInjection, err := di.InitializeDI()
	r := dependencyInjection.Router
	r.Route("/id", func(r chi.Router) {
		r.Post("/", dependencyInjection.IdController.PostId)
		r.Get("/{id}", dependencyInjection.IdController.GetId)
	})
	err = http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")), r)
	if err != nil {
		log.Fatalf("Err http server: %v", err)
	}
}
