package store

import (
    "github.com/sidekick-coder/atlas/internal/models"
)

func (s *Store) GetEntryMetasByEntryID(entryID int64) ([]*models.EntryMeta, error) {
	smtmt := `
	SELECT entry_id, name, value
	FROM entry_metas
	WHERE entry_id = $1;
	`
	rows, err := s.db.Query(smtmt, entryID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var metas []*models.EntryMeta

	for rows.Next() {
		var meta models.EntryMeta

		err := rows.Scan(&meta.EntryID, &meta.Name, &meta.Value)

		if err != nil {
			return nil, err
		}

		metas = append(metas, &meta)
	}

	return metas, nil
}

func (s *Store) UpdateEntryMeta(u *models.EntryMeta) error {
    // UPDATE ...
    return nil
}

func (s *Store) UpsertEntryMeta(e *models.EntryMeta) error {
	smtmt := `
	INSERT INTO entry_metas (entry_id, name, value)
	VALUES ($1, $2, $3)
	ON CONFLICT (entry_id, name) DO UPDATE SET value = EXCLUDED.value;
	`

	_, err := s.db.Exec(smtmt, e.EntryID, e.Name, e.Value)

	if err != nil {
		return err
	}

    return nil
}
