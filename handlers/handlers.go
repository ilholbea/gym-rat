package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ilholbea/gym-rat/services"
)

type AccountHandlers struct {
	service services.AccountService
}

func NewAccountHandlers(service services.AccountService) *AccountHandlers {
	return &AccountHandlers{service: service}
}

// Health handles HTTP GET requests for checking the health status of the service.
func (h *AccountHandlers) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// GetAccountsHandler handles HTTP GET requests for retrieving a list of all accounts.
func (h *AccountHandlers) GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	accounts, err := h.service.GetAll(ctx)
	if err != nil {
		log.Errorf("Error getting accounts: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if accounts == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	accountJSON, err := json.Marshal(accounts)
	if err != nil {
		log.Errorf("Error marshalling accounts: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(accountJSON)
	if err != nil {
		log.Errorf("Error writing response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// CreateAccountHandler handles HTTP POST requests for creating a new account.
func (h *AccountHandlers) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

// UpdateAccountHandler handles HTTP PUT requests for updating an existing account.
func (h *AccountHandlers) UpdateAccountHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

// GetAccountHandler handles HTTP GET requests for retrieving a specific account by its ID.
func (h *AccountHandlers) GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

// DeleteAccountHandler handles HTTP DELETE requests for deleting an existing account by its ID.
func (h *AccountHandlers) DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
