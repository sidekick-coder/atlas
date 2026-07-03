package metadata

import (
	"github.com/sidekick-coder/atlas/internal/models"
)

func Unset(info *models.EntryInfo, name string, handlers []Handler) (bool, error) {

	for _, handler := range handlers {
		err := handler.Unset(info, name)

		if err != nil {
			return false, err
		}
	}

	return true, nil
}
