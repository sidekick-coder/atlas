
package db

import (
	"database/sql"
)

func DeleteEntry(db *sql.DB, path string) error {
	sql := "DELETE FROM entries WHERE path = ?"

	_, err := db.Exec(sql, path)

	return err
}
