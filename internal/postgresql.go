package internal

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/ilholbea/gym-rat/config"
	"github.com/ilholbea/gym-rat/types"
	_ "github.com/lib/pq"
)

type Postgres struct {
	database *sql.DB
}

func New(c *config.Database) (Postgres, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return Postgres{}, err
	}

	return Postgres{database: db}, err
}

func (pg *Postgres) Create(exercise *types.Exercise) error {
	exercise.ID = generateUUID()
	insertStatement := `INSERT INTO "gym-rat".exercise(id, name, description, video, image) VALUES($1, $2, $3, $4, $5);`

	_, err := pg.database.Exec(insertStatement, exercise.ID, exercise.Name, exercise.Description, exercise.Video, exercise.Image)
	if err != nil {
		return err
	}
	return nil
}

func (pg *Postgres) GetAll() ([]types.Exercise, error) {
	var response []types.Exercise

	selectStatement := `SELECT * from "gym-rat".exercise`
	rows, err := pg.database.Query(selectStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var retrievedExercise types.Exercise
		if err = rows.Scan(&retrievedExercise.ID, &retrievedExercise.Name, &retrievedExercise.Description, &retrievedExercise.Video, &retrievedExercise.Image); err != nil {
			return nil, err
		}
		response = append(response, retrievedExercise)
	}
	return response, nil
}

func (pg *Postgres) Get(id string) (*types.Exercise, error) {
	var response []types.Exercise
	selectStatement := `SELECT * from "gym-rat".exercise where id=$1`
	rows, err := pg.database.Query(selectStatement, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var retrievedExercise types.Exercise
		if err = rows.Scan(&retrievedExercise.ID, &retrievedExercise.Name, &retrievedExercise.Description, &retrievedExercise.Video, &retrievedExercise.Image); err != nil {
			return nil, err
		}
		response = append(response, retrievedExercise)
	}

	if len(response) > 1 {
		return nil, fmt.Errorf("too many number of results with the id: %s", id)
	}

	if len(response) == 0 {
		return nil, fmt.Errorf("no results found with the id: %s", id)
	}

	return &response[0], nil
}

func (pg *Postgres) Delete(id string) error {
	found, err := pg.Get(id)
	if err != nil {
		return err
	}

	if found == nil {
		return fmt.Errorf("no results found with the id: %s", id)
	}

	deleteStatement := `DELETE from "gym-rat".exercise where id=$1`
	_, err = pg.database.Exec(deleteStatement, id)

	if err != nil {
		return err
	}
	return nil
}

func (pg *Postgres) Update(exercise *types.Exercise) error {
	found, err := pg.Get(exercise.ID)
	if err != nil {
		return err
	}

	if found == nil {
		return fmt.Errorf("no results found with the id: %s", exercise.ID)
	}

	updateStatement := `UPDATE "gym-rat".exercise SET name=$1, description=$2, video=$3, image=$4 where id=$5`
	_, err = pg.database.Exec(updateStatement, exercise.Name, exercise.Description, exercise.Video, exercise.Image, exercise.ID)
	if err != nil {
		return err
	}

	return nil
}

func generateUUID() string {
	generatedUUID := uuid.New()
	return generatedUUID.String()
}
