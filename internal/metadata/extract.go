package metadata

import (
	"maps"
	"strings"

	"github.com/sidekick-coder/atlas/internal/models"
)

func Extract(info *models.EntryInfo, handlers []Handler) (map[string]string, error) {
	metadata := make(map[string]string)

	ids := []string{}

	for _, handler := range handlers {
		data, err := handler.Extract(info)

		ids = append(ids, handler.ID())

		if err != nil {
			return nil, err
		}

		maps.Copy(metadata, data)
	}

	metadata["basename"] = info.BaseName
	metadata["type"] = info.Type
	metadata["ext"] = strings.TrimPrefix(info.Ext, ".")

	return metadata, nil
}
