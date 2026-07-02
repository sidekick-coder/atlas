package entry

import (
	"github.com/sidekick-coder/atlas/internal/models"
)

func (r * Repository) List() ([]models.Entry, error) {
	// SELECT ...
	smtmt := `
	SELECT id, path, is_dir
	FROM entries;
	`

    rows, err := r.Database.Query(smtmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var entries []models.Entry

	for rows.Next() {
		var entry models.Entry

		err := rows.Scan(&entry.ID, &entry.Path, &entry.IsDir)

		if err != nil {
			return nil, err
		}

		entries = append(entries, entry)
	}

	return entries, nil
}
