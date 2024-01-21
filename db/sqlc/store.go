package db

import "database/sql"

type Store struct {
	db *sql.DB
	*Queries
}
