package metadata

import (
	"github.com/sidekick-coder/atlas/internal/models"
)

func Set(info *models.EntryInfo, name string, value string, handlers []Handler) (bool, error) {

	for _, handler := range handlers {
		updated, err := handler.Set(info, name, value)

		if err != nil {
			return false, err
		}

		if updated {
			return true, nil
		}
	}

	return false, nil
}
