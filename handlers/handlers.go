package handlers

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io/ioutil"
	"log"
	"net/http"
)

func createExercise(w http.ResponseWriter, r *http.Request) {
	newExercise := exercise{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		log.Fatalf("unable to read request body, %v", err)
	}

	err = json.Unmarshal(body, &newExercise)
	if err != nil {
		w.WriteHeader(500)
		log.Fatalf("unable to unmarshal, %v", err)
	}

	db, err := connectToDB()
	if err != nil {
		w.WriteHeader(500)
		log.Fatalf("unable to connect to DB, %v", err)
	}

	insertStatement := `INSERT INTO exercise(id, name, description, video, image) VALUES(` + newExercise.ID + ",'" + newExercise.Name + "','" + newExercise.Description + "','" + newExercise.Video + "','" + newExercise.Image + "')"

	_, err = db.Exec(insertStatement)
	if err != nil {
		w.WriteHeader(500)
		log.Fatalf("unable to execute query, %v", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
}

func getExercises(w http.ResponseWriter, r *http.Request) {
	var response []exercise
	db, err := connectToDB()
	if err != nil {
		w.WriteHeader(500)
		log.Fatalf("unable to connect to DB, %v", err)
	}

	selectStatement := `SELECT * from exercise`

	rows, err := db.Query(selectStatement)
	if err != nil {
		w.WriteHeader(500)
		log.Fatalf("unable to execute query, %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var ex exercise
		if err := rows.Scan(&ex.ID, &ex.Name, &ex.Description, &ex.Video, &ex.Image); err != nil {
			return
		}
		response = append(response, ex)
	}

	resp, err := json.Marshal(response)
	if err != nil {
		log.Fatal("unable to marshal response")
	}

	_, err = w.Write(resp)
	if err != nil {
		log.Fatal("unable to write response")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
}

func main() {
	router := chi.NewRouter()

	router.Get("/exercise", getExercises)
	router.Post("/exercise", createExercise)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("unable to start server")
	}

}
