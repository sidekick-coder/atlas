package db

import (
	"database/sql"
)

func UpsertEntry(db *sql.DB, path string, isDir bool) error {
	_, err := db.Exec(`
		INSERT OR REPLACE INTO entries (path, is_dir)
		VALUES (?, ?)
	`, path, isDir)

	return err
}
