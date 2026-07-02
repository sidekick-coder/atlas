package entry

import (
	"github.com/sidekick-coder/atlas/internal/models"
)

func (r * Repository) GetByPath(path string) (*models.Entry, error) {
	// SELECT ...
	smtmt := `
	SELECT id, path
	FROM entries
	WHERE path = $1
	LIMIT 1;
	`

    rows, err := r.Database.Query(smtmt, path)

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

    return nil, nil
}
