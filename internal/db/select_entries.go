package db

import "database/sql"

type Entry struct {
	Path  string
	IsDir bool
}

func SelectEntries(db *sql.DB) ([]Entry, error) {
	rows, err := db.Query(`
		SELECT path, is_dir
		FROM entries
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []Entry

	for rows.Next() {
		var e Entry

		if err := rows.Scan(&e.Path, &e.IsDir); err != nil {
			return nil, err
		}

		entries = append(entries, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}
