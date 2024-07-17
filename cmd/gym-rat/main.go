package main

import (
	"context"
	"fmt"
	"log"
	http "net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"

	"github.com/ilholbea/gym-rat/config"
	"github.com/ilholbea/gym-rat/database"
	"github.com/ilholbea/gym-rat/handlers"
	"github.com/ilholbea/gym-rat/repository"
	"github.com/ilholbea/gym-rat/routes"
	"github.com/ilholbea/gym-rat/services"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("unable to find env file")
	}

	conf, err := config.NewConfig()
	if err != nil {
		panic("unable to load db config")
	}

	// Connect to the db
	db, err := database.ConnectDB(conf.DatabaseConfig)
	if err != nil {
		log.Fatalf("Could not connect to the db: %v\n", err)
	}
	defer db.Close()

	// Init repositories
	accountRepo := repository.NewAccountRepository(db)

	// Init services
	accountService := services.NewAccountService(accountRepo)

	// Init handlers
	accountHandlers := handlers.NewAccountHandlers(accountService)

	// Init router
	router := routes.NewRouter(accountHandlers)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on port ':8080': %s\n", err)
		}
	}()
	fmt.Println("Server is ready to handle requests at :8080")

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}
	log.Println("Server stopped")
}
