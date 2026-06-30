package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func Connect(dbPath string) (*sql.DB, error) {
	d, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	if err := d.Ping(); err != nil {
		return nil, err
	}

	return d, nil
}
