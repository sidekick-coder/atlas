package database

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Connect(filename string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite", filename)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}


