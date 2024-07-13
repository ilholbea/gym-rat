package handlers

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Server struct {
	Router   *chi.Mux
	Database *sql.DB
}

func NewServer(router *chi.Mux, db *sql.DB) (*Server, error) {

	return &Server{Router: router, Database: db}, nil
}

func (s *Server) Health(w http.ResponseWriter, _ *http.Request) {
	log.Printf("health - endpoint was accessed")
	w.WriteHeader(http.StatusOK)
}
