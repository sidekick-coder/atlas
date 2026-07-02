package entrymeta

import (
	"github.com/sidekick-coder/atlas/internal/models"
)

func (r * Repository) UpsertMany(entryId int64, metas map[string]string) ([]models.EntryMeta, error) {
	smtmt := `
	INSERT INTO entry_metas (entry_id, name, value)
	VALUES ($1, $2, $3)
	ON CONFLICT (entry_id, name) DO UPDATE SET value = EXCLUDED.value;
	`

	for name, value := range metas {
		_, err := r.Database.Query(smtmt, entryId, name, value)

		if err != nil {
			return nil, err
		}
	}

	return r.ListByEntryID(entryId)
}
