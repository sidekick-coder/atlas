package database 

import (
	"database/sql"
)

func (db *Database) Exec(query string, args ...any) (sql.Result, error) {
	return db.Connection.Exec(query, args...)
}
