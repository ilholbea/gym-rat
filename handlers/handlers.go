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

var dbConfig = &config.Database{
	Host:     "localhost",
	Port:     "5432",
	User:     "postgres",
	Password: "PASSWORD1!",
	Name:     "gym-rat",
}

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

	db, err := internal.New(dbConfig)

	if err != nil {
		log.Errorf("internal server error: unable to connect to database: %v", err)
		w.WriteHeader(500)
		return
	}

	err = db.Create(&newExercise)
	if err != nil {
		log.Errorf("internal server error: unable to create new exercise: %v", err)
		w.WriteHeader(500)
		return
	}
}

func updateExercise(w http.ResponseWriter, r *http.Request) {
	updatedExercise := types.Exercise{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		log.Errorf("internal server error: unable to read request body, %v", err)
	}

	err = json.Unmarshal(body, &updatedExercise)
	if err != nil {
		log.Errorf("internal server error: unable to unmarshal body: %v", err)
		w.WriteHeader(500)
		return
	}

	db, err := internal.New(dbConfig)

	if err != nil {
		log.Errorf("internal server error: unable to connect to database: %v", err)
		w.WriteHeader(500)
		return
	}

	err = db.Update(&updatedExercise)
	if err != nil {
		log.Errorf("internal server error: unable to create new exercise: %v", err)
		w.WriteHeader(500)
		return
	}
}

func getExercises(w http.ResponseWriter, r *http.Request) {
	db, err := internal.New(dbConfig)

	if err != nil {
		log.Errorf("internal server error: unable to connect to database: %v", err)
		w.WriteHeader(500)
		return
	}

	exercises, err := db.GetAll()
	if err != nil {
		log.Errorf("internal server error: unable to retrieve list of exercises: %v", err)
		w.WriteHeader(500)
		return
	}

	response, err := json.Marshal(exercises)
	if err != nil {
		log.Errorf("internal server error: unable to marshal response: %v", err)
		w.WriteHeader(500)
		return
	}

	_, err = w.Write(response)
	if err != nil {
		log.Errorf("internal server error: unable to write response: %v", err)
		w.WriteHeader(500)
		return
	}
}

func getExercise(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		log.Errorf("bad request: missing parameter")
		w.WriteHeader(400)
		return
	}

	db, err := internal.New(dbConfig)

	if err != nil {
		log.Errorf("internal server error: unable to connect to database: %v", err)
		w.WriteHeader(500)
		return
	}

	exercise, err := db.Get(id)
	if err != nil {
		log.Errorf("internal server error: unable to retrieve exercise: %v", err)
		w.WriteHeader(500)
		return
	}

	response, err := json.Marshal(exercise)
	if err != nil {
		log.Errorf("internal server error: unable to marshal response: %v", err)
		w.WriteHeader(500)
		return
	}

	_, err = w.Write(response)
	if err != nil {
		log.Errorf("internal server error: unable to write response: %v", err)
		w.WriteHeader(500)
		return
	}
}

func deleteExercise(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		log.Errorf("bad request: missing parameter")
		w.WriteHeader(400)
		return
	}

	db, err := internal.New(dbConfig)

	if err != nil {
		log.Errorf("internal server error: unable to connect to database: %v", err)
		w.WriteHeader(500)
		return
	}

	err = db.Delete(id)
	if err != nil {
		log.Errorf("internal server error: unable to delete exercise: %v", err)
		w.WriteHeader(500)
		return
	}
}

func Router() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/exercise", createExercise)
	router.Get("/exercise", getExercises)
	router.Get("/exercise/{id}", getExercise)
	router.Delete("/exercise/{id}", deleteExercise)
	router.Put("/exercise", updateExercise)

	return router
}
