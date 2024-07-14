package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/ilholbea/gym-rat/handlers"
	"net/http"
)

func NewRouter(accountHandlers *handlers.AccountHandlers) http.Handler {
	router := chi.NewRouter()

	// Account Routes
	router.Get("/health", accountHandlers.Health)

	return router
}
