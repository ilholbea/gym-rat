package database

import (
	"database/sql"
	"fmt"
	"github.com/ilholbea/gym-rat/config"
	_ "github.com/lib/pq"
)

// ConnectDB establishes a connection to the PostgreSQL database
func ConnectDB(config *config.DatabaseConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}
