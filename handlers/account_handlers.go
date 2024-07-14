package handlers

import (
	"github.com/ilholbea/gym-rat/services"
	"net/http"
)

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
