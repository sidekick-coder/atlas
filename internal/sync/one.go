package sync

import (
	"database/sql"
	"path/filepath"
	"fmt"

	"github.com/sidekick-coder/atlas/internal/models"
	"github.com/sidekick-coder/atlas/internal/drive"
	"github.com/sidekick-coder/atlas/internal/extract"
	"github.com/sidekick-coder/atlas/internal/store"
)


func One(conn *sql.DB, root string, path string, is_dir bool) error {
	store := store.New(conn)

	entry, err := store.UpsertEntry(path, is_dir)


	if err != nil {
		return err
	}

	meta := map[string]any{}

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

	for k, v := range meta {
		model := models.EntryMeta{
			Name:  k,
			EntryID: entry.ID,
			Value: fmt.Sprintf("%v", v),
		}

		store.UpsertEntryMeta(&model)
	}

	return nil
}

