package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/ilholbea/gym-rat/config"
	"github.com/ilholbea/gym-rat/internal"
	"github.com/ilholbea/gym-rat/types"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

const (
	Host     = "localhost"
	Port     = "5432"
	User     = "postgres"
	Password = "-"
	Name     = "gymRat"
)

func createExercise(w http.ResponseWriter, r *http.Request) {
	newExercise := types.Exercise{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		log.Errorf("internal server error: unable to read request body, %v", err)
	}

	err = json.Unmarshal(body, &newExercise)
	if err != nil {
		log.Errorf("internal server error: unable to unmarshal body: %v", err)
		w.WriteHeader(500)
		return
	}

	db, err := internal.New(&config.Database{
		Host:     Host,
		Port:     Port,
		User:     User,
		Password: Password,
		Name:     Name,
	})

	err = db.Create(&newExercise)
	if err != nil {
		log.Errorf("internal server error: unable to create new exercise: %v", err)
		w.WriteHeader(500)
		return
	}
}

func Router() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/exercise", createExercise)

	return router
}
