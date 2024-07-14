package handlers

import (
	"github.com/ilholbea/gym-rat/services"
	"net/http"
)

//type Server struct {
//	Router   *chi.Mux
//	Database *sql.DB
//}
//
//func NewServer(router *chi.Mux, db *sql.DB) (*Server, error) {
//
//	return &Server{Router: router, Database: db}, nil
//}
//
//func (s *Server) Health(w http.ResponseWriter, _ *http.Request) {
//	log.Printf("health - endpoint was accessed")
//	w.WriteHeader(http.StatusOK)
//}

type AccountHandlers struct {
	service services.AccountService
}

func NewAccountHandlers(service services.AccountService) *AccountHandlers {
	return &AccountHandlers{service: service}
}

func (h *AccountHandlers) Health(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (h *AccountHandlers) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
