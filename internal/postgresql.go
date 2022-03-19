package internal

import (
	"database/sql"
	"fmt"
	"github.com/ilholbea/gym-rat/config"
	"github.com/ilholbea/gym-rat/types"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
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
	insertStatement := `INSERT INTO exercise(name, description, video, image) VALUES($1, $2, $3, $4) returning uuid`

	var uuid string
	err := pg.database.QueryRow(insertStatement, exercise.Name, exercise.Description, exercise.Video, exercise.Image).Scan(&uuid)
	if err != nil {
		return err
	}

	log.Printf("Inserted record with id '%s'", uuid)
	return nil
}
