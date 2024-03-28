package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/ilholbea/gym-rat/config"
	"log"
	"net/http"
)

type Server struct {
	Router   *chi.Mux
	Database config.Database
}

func NewServer(router *chi.Mux, conf *config.Config) (*Server, error) {
	router.Get("/health", health)
	return &Server{Router: router, Database: conf.DbConfig}, nil
}

func health(w http.ResponseWriter, _ *http.Request) {
	log.Printf("health - endpoint was accessed")
	w.WriteHeader(http.StatusOK)
}
