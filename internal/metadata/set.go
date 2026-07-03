package metadata

import (
	"github.com/sidekick-coder/atlas/internal/drive/v2"
)

func Set(info *drive.EntryInfo, name string, value string, handlers []Handler) (bool, error) {

	for _, handler := range handlers {
		err := handler.Set(info, name, value)

		if err != nil {
			return false, err
		}
	}

	return true, nil
}
