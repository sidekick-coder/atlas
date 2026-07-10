package entry

import (
	"strings"

	"github.com/sidekick-coder/atlas/internal/models"
)

func (r *Repository) GetEntryMetasMap(entryMetas []models.EntryMeta, entryId int64) map[string]string {
	result := make(map[string]string)

	for _, meta := range entryMetas {
		if meta.EntryID == entryId {
			result[meta.Name] = meta.Value
		}
	}

	return result
}

func (r *Repository) ListMetas(entries ...models.Entry) ([]models.EntryMeta, error) {
	// SELECT ...
	smtmt := []string{}
	params := []any{}

	smtmt = append(smtmt, "SELECT id, entry_id, name, value")
	smtmt = append(smtmt, "FROM entry_metas")

	smtmt = append(smtmt, "WHERE entry_id in (" + strings.Repeat("?,", len(entries)-1) + "?)")
	
	for _, entry := range entries {
		params = append(params, entry.ID)
	}

	rows, err := r.Database.Query(strings.Join(smtmt, " "), params...)

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
