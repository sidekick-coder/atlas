package sync

import (
	"database/sql"

	"github.com/sidekick-coder/atlas/internal/db"
	"github.com/sidekick-coder/atlas/internal/drive"
)


func All(conn *sql.DB, root string) error {
	return drive.ScanStream(root, func(e drive.Entry) error {
		return db.UpsertEntry(conn, e.Path, e.IsDir)
	})
}
