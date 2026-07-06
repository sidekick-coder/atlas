package database 

import (
	"database/sql"
)

func (db *Database) Query(query string, args ...any) (*sql.Rows, error) {
	rows, err := db.Connection.Query(query, args...)

	if err != nil {
		return nil, err
	}

	return rows, nil
}
