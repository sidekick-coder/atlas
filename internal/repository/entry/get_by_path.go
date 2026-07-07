package entry

import (
	"fmt"

	"github.com/sidekick-coder/atlas/internal/models"
)

func (r *Repository) GetByPath(path string) (*models.Entry, error) {
	smtmt := "SELECT id, path FROM entries WHERE path = $1 LIMIT 1"
	params := []any{path}

	rows, err := r.Database.Query(smtmt, params...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		var entry models.Entry

		err := rows.Scan(&entry.ID, &entry.Path)

		if err != nil {
			return nil, err
		}

		return &entry, nil
	}

	return nil, fmt.Errorf("entry not found for path: %s", path)
}
