package sync

import (
	"database/sql"

	"github.com/sidekick-coder/atlas/internal/db"
	"github.com/sidekick-coder/atlas/internal/drive"
)

func Deleted(conn *sql.DB, seen map[string]struct{}) error {
	dbEntries, err := db.SelectEntries(conn)

	if err != nil {
		return err
	}

	for _, e := range dbEntries {
		if _, ok := seen[e.Path]; !ok {
			db.DeleteEntry(conn, e.Path)
		}
	}

	return nil
}

func Files(conn *sql.DB, root string) (map[string]struct{}, error) {
	seen := make(map[string]struct{}) 

	err := drive.ScanStream(root, func(e drive.Entry) error {
		seen[e.Path] = struct{}{}

		return db.UpsertEntry(conn, e.Path, e.IsDir)
	})

	return seen, err
}

func All(conn *sql.DB, root string) error {
	seen, err := Files(conn, root)

	if err != nil {
		return err
	}

	err = Deleted(conn, seen)

	if err != nil {
		return err
	}

	return nil
}
