package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ilholbea/gym-rat/handlers"
)

func NewRouter(accountHandlers *handlers.AccountHandlers) http.Handler {
	router := chi.NewRouter()

	router.Get("/health", accountHandlers.Health)

	// Account Routes
	router.Get("/accounts", accountHandlers.GetAccountsHandler)
	router.Post("/accounts", accountHandlers.CreateAccountHandler)
	router.Get("/accounts/{accountId}", accountHandlers.GetAccountHandler)
	router.Put("/accounts/{accountId}", accountHandlers.UpdateAccountHandler)
	router.Delete("/accounts/{accountId}", accountHandlers.DeleteAccountHandler)

	return router
}
