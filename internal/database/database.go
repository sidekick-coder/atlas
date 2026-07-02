package database

import (
	"database/sql"
)

type Database struct {
	Connection *sql.DB
}

func New(conn *sql.DB) *Database {
	return &Database{Connection: conn}
}

