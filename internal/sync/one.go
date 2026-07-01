package sync

import (
	"database/sql"
	"path/filepath"
	"fmt"

	"github.com/sidekick-coder/atlas/internal/db"
	"github.com/sidekick-coder/atlas/internal/drive"
	"github.com/sidekick-coder/atlas/internal/extract"
)


func One(conn *sql.DB, root string, path string, is_dir bool) error {
	err := db.UpsertEntry(conn, path, is_dir)
	meta := map[string]any{}

	if err != nil {
		return err
	}

	isMakdown, err := filepath.Match("*.md", filepath.Base(path))

	if err != nil {
		return err
	}

	if isMakdown {
		err, content := drive.ReadAsString(root, path)

		if err != nil {
			return err
		}

		extract.Markdown(content, &meta)
	}

	fmt.Printf("path: %s, is_dir: %v\n", path, is_dir)
	for k, v := range meta {
		fmt.Printf("--%s = %v\n", k, v)
	}



	return nil
}

