package metadata

import (
	"github.com/sidekick-coder/atlas/internal/models"
)

func Set(info *models.EntryInfo, name string, value string, handlers []Handler) (bool, error) {
	m, err := Create(info)

	if err != nil {
		return false, err
	}

	return m.Set(name, value)
}
