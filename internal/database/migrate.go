package database 

import (
	"os"
	"path/filepath"
	"time"
)

func (db *Database) Migrate() error {
	if _, err := db.Connection.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE,
			applied_at INTEGER
		);
	`); err != nil {
		return err
	}

	dir := "internal/database/migrations"

	files, err := os.ReadDir(dir)

	if err != nil {
		return err
	}

	for _, f := range files {
		name := f.Name()

		var exists int

		err := db.Connection.QueryRow(
			"SELECT 1 FROM migrations WHERE name = ?",
			name,
		).Scan(&exists)

		if err == nil {
			continue // already applied
		}

		path := filepath.Join(dir, name)
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		if _, err := db.Connection.Exec(string(content)); err != nil {
			return err
		}

		_, err = db.Connection.Exec(
			"INSERT INTO migrations (name, applied_at) VALUES (?, ?)",
			name,
			time.Now().Unix(),
		)
		if err != nil {
			return err
		}
	}

	return nil
}

