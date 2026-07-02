package entrymeta

import (
	"github.com/sidekick-coder/atlas/internal/models"
)

func (r * Repository) ListByEntryId(entryId int64) ([]models.EntryMeta, error) {
	// SELECT ...
	smtmt := `
	SELECT id, entry_id, name, value
	FROM entry_metas
	WHERE id = $1;
	`

    rows, err := r.Database.Query(smtmt, entryId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var metas []models.EntryMeta

	for rows.Next() {
		var meta models.EntryMeta

		err := rows.Scan(&meta.ID, &meta.EntryID, &meta.Name, &meta.Value)

		if err != nil {
			return nil, err
		}

		metas = append(metas, meta)
	}

	return metas, nil
}
