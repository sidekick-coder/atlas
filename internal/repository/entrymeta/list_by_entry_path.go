package entrymeta

import (
	"github.com/sidekick-coder/atlas/internal/models"
)

func (r * Repository) ListByEntryPath(path string) ([]models.EntryMeta, error) {
	entry, err := r.EntryRepository.GetByPath(path)

	if err != nil {
		return nil, err
	}

	return r.ListByEntryID(entry.ID)
}
