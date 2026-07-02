package store

import (
    "github.com/sidekick-coder/atlas/internal/models"
)

func (s *Store) CreateEntry(u *models.Entry) error {
    // INSERT ...
    return nil
}

func (s *Store) GetEntry(id int64) (*models.Entry, error) {
    // SELECT ...
    return nil, nil
}

func (s *Store) GetEntryByPath(path string) (*models.Entry, error) {
	// SELECT ...
	smtmt := `
	SELECT id, path, is_dir
	FROM entries
	WHERE path = $1
	LIMIT 1;
	`

    rows, err := s.db.Query(smtmt, path)

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

func (s *Store) UpdateEntry(u *models.Entry) error {
    // UPDATE ...
    return nil
}

func (s *Store) UpsertEntry(path string, is_dir bool) (*models.Entry, error) {
	smtmt := `
	INSERT INTO entries (path, is_dir)
	VALUES ($1, $2)
	ON CONFLICT (path) DO UPDATE SET is_dir = EXCLUDED.is_dir;
	`

	_, err := s.db.Exec(smtmt, path, is_dir)

	if err != nil {
		return nil, err
	}

	return s.GetEntryByPath(path)
}

func (s *Store) DeleteEntry(id int64) error {
    // DELETE ...
    return nil
}
