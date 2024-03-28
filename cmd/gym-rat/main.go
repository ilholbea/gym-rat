package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/ilholbea/gym-rat/config"
	"github.com/ilholbea/gym-rat/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("unable to find env file")
	}

	conf, err := config.NewConfig()
	if err != nil {
		panic("unable to load config")
	}

	router := chi.NewRouter()
	server, err := handlers.NewServer(router, conf)
	if err != nil {
		panic("unable to start server")
	}

	log.Printf("Server Started...")

	// TODO: Make port configurable
	err = http.ListenAndServe(":8080", server.Router)
	if err != nil {
		log.Fatalf("Unable to start server: %v", err)
	}
}
