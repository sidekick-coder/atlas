package metadata

import (
	"github.com/sidekick-coder/atlas/internal/models"
)

func Unset(info *models.EntryInfo, name string, handlers []Handler) (bool, error) {
	m, err := Create(info)

	if err != nil {
		return false, err
	}

	return m.Unset(name)
}
