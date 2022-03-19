package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

)

const (
	host     = "127.0.0.1"
	port     = "5432"
	user     = "postgres"
	password = "Microx23!"
	dbname   = "gymRat"
)

type postgreDB struct
}
func (pg *postgreDB) connectToDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
