package service

import "database/sql"

type database interface {
	connectToDatabase(error, *sql.DB)
}
